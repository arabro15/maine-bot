package user

import (
	"errors"
	"maineBot/domain/entity"
	"maineBot/domain/entity/book"
	"time"
)

// Errors
var (
	ErrEmptyUserID         = errors.New("empty UserID")
	ErrEmptyName           = errors.New("empty Name")
	ErrEmptyTelegramName   = errors.New("empty TelegramName")
	ErrEmptyBook           = errors.New("empty Book")
	ErrCheckRequiredFields = errors.New("func CheckRequiredFields is error")
)

type User struct {
	id           ID
	name         entity.Name
	telegramName string
	age          Age
	gender       Gender
	book         book.Book
	createdAt    time.Time
	updatedAt    time.Time
	status       Status
}

func (u User) NewUser(
	id ID,
	name entity.Name,
	telegramName string,
	book book.Book,
	status Status,
) (*User, error) {

	check := u.checkRequiredFields(&id, &name, &telegramName, &book)
	if check != nil {
		return nil, ErrCheckRequiredFields
	}

	return &User{
		id:           id,
		name:         name,
		telegramName: telegramName,
		book:         book,
		status:       status,
	}, nil
}

func (u User) Id() ID {
	return u.id
}

func (u User) Name() entity.Name {
	return u.name
}

func (u User) TelegramName() string {
	return u.telegramName
}

func (u User) Age() Age {
	return u.age
}

func (u User) Gender() Gender {
	return u.gender
}

func (u User) Book() book.Book {
	return u.book
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u User) Status() Status {
	return u.status
}

func (u User) checkRequiredFields(
	id *ID,
	name *entity.Name,
	telegramName *string,
	book *book.Book,
) error {
	if len(id.Value().String()) == 0 {
		return ErrEmptyUserID
	}
	if len(name.Value()) == 0 {
		return ErrEmptyName
	}
	if len(*telegramName) == 0 {
		return ErrEmptyTelegramName
	}
	if len(book.BookID().Value().String()) == 0 {
		return ErrEmptyBook
	}

	return nil
}
