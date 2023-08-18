package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"maineBot/adapter/repository/converter"
	"maineBot/adapter/repository/model"
	"maineBot/boundary/repository"
	"maineBot/domain/entity/user"
	"maineBot/pkg/client/postgresql"
)

type userRepositoryImpl struct {
	client postgresql.Client
}

func (ur *userRepositoryImpl) InsertUser(ctx context.Context, user user.User) (userID user.ID, err error) {
	q := `
		INSERT INTO "user" 
		    (
		     id,
		     name, 
		     telegram_name, 
		     age, 
		     gender, 
		     created_at, 
		     updated_at, 
		     status
		    ) 
		VALUES 
		    (
		     $id, 
		     $name, 
		     $telegramName, 
		     $age, $gender, $createdAt, $updatedAt, $status)
		RETURNING id
	`
	userDbModel := converter.ToModel(&user)
	err = ur.client.QueryRow(ctx, q, userDbModel).Scan(&userID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			return userID, newErr
		}
		return userID, err
	}

	return userID, nil
}

func (ur *userRepositoryImpl) DeleteUserByID(ctx context.Context, userID user.ID) (bool, error) {
	q := `
		DELETE FROM "user"
		WHERE id := $id
	`

	commandTag, err := ur.client.Exec(ctx, q, userID.String())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			return commandTag.Delete(), newErr
		}

		return commandTag.Delete(), err
	}

	return commandTag.Delete(), nil
}

func (ur *userRepositoryImpl) UpdateUser(ctx context.Context, user user.User) (bool, error) {
	q := `
		UPDATE "user" 
		SET  id = $id,
		     name = $name, 
		     telegram_name = $telegramName, 
		     age = $age, 
		     gender = $gender, 
		     created_at = $createdAt, 
		     updated_at = $updatedAt, 
		     status = $status
		WHERE id = $id
	`
	userDbModel := converter.ToModel(&user)
	commandTag, err := ur.client.Exec(ctx, q, userDbModel)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			return commandTag.Update(), newErr
		}
		return commandTag.Update(), err
	}

	return commandTag.Update(), nil
}

func (ur *userRepositoryImpl) FindUserByID(ctx context.Context, id user.ID) (user.User, error) {
	q := `
		SELECT
		     id,
		     name, 
		     telegram_name, 
		     age, 
		     gender, 
		     created_at, 
		     updated_at, 
		     status
		FROM "user"
		WHERE id
	`
	var userDbModel *model.UserDbModel
	user, errConverter := converter.ToEntity(*userDbModel)
	if errConverter != nil {
		return *user, errors.New("userDbModel can`t converter to user")
	}

	if err := ur.client.QueryRow(ctx, q, id.Value()).Scan(userDbModel); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			return *user, newErr
		}
		return *user, err
	}

	return *user, nil
}

func (ur *userRepositoryImpl) FindUserByTelegramName(ctx context.Context, telegramName string) (user.User, error) {
	q := `
		SELECT
		     id,
		     name, 
		     telegram_name, 
		     age, 
		     gender, 
		     created_at, 
		     updated_at, 
		     status
		FROM public."user"
		WHERE telegram_name = $telegramName
	`
	var userDbModel *model.UserDbModel
	user, errConverter := converter.ToEntity(*userDbModel)
	if errConverter != nil {
		return *user, errors.New("userDbModel can`t converter to user")
	}

	if err := ur.client.QueryRow(ctx, q, telegramName).Scan(userDbModel); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			return *user, newErr
		}
		return *user, err
	}

	return *user, nil
}

func (ur *userRepositoryImpl) FindAllUser(ctx context.Context) ([]user.User, error) {
	q := `
		SELECT 
		    id,
		    name,
		    telegram_name,
		    age, 
		    gender,
		    created_at, 
		    updated_at, 
		    status   
		FROM public."user";
	`
	rows, err := ur.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]user.User, 0)

	for rows.Next() {
		var userDbModel model.UserDbModel
		err = rows.Scan(userDbModel)
		if err != nil {
			return nil, err
		}
		u2, errConverter := converter.ToEntity(userDbModel)
		if errConverter != nil {
			return nil, errors.New("userDbModel can`t converter to user")
		}

		users = append(users, *u2)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserRepository(client postgresql.Client) repository.UserRepository {
	return &userRepositoryImpl{
		client: client,
	}
}
