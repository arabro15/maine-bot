package user

import (
	"errors"
	"maineBot/domain/entity"
	"maineBot/domain/entity/book"
	"time"
)

// Errors
var (
	ErrCheckRequiredFieldsBuilder = errors.New("check required fields in Builder is Error")
	ErrEmptyUserIdBuilder         = errors.New("empty UserID in Builder")
	ErrEmptyNameBuilder           = errors.New("empty Name in Builder")
	ErrEmptyTelegramNameBuilder   = errors.New("empty TelegramName in Builder")
	ErrEmptyBookBuilder           = errors.New("empty Book in Builder")
	ErrIsZeroTimeCreatedAtBuilder = errors.New("time CreatedAt is Zero in Builder")
	ErrIsZeroTimeUpdatedAtBuilder = errors.New("time UpdatedAt is Zero in Builder")
)

type Builder interface {
	UserID(id ID) Builder
	Name(name entity.Name) Builder
	TelegramName(telegramName string) Builder
	Age(age Age) Builder
	Gender(gender Gender) Builder
	Book(book book.Book) Builder
	CreatedAt(createdAt time.Time) Builder
	UpdatedAt(UpdatedAt time.Time) Builder
	UserStatus(status Status) Builder
	build() (*User, error)
}

type userBuilder struct {
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

func (builder *userBuilder) UserID(id ID) Builder {
	builder.id = id
	return builder
}

func (builder *userBuilder) Name(name entity.Name) Builder {
	builder.name = name
	return builder
}

func (builder *userBuilder) TelegramName(telegramName string) Builder {
	builder.telegramName = telegramName
	return builder
}

func (builder *userBuilder) Age(age Age) Builder {
	builder.age = age
	return builder
}

func (builder *userBuilder) Gender(gender Gender) Builder {
	builder.gender = gender
	return builder
}

func (builder *userBuilder) Book(book book.Book) Builder {
	builder.book = book
	return builder
}

func (builder *userBuilder) CreatedAt(createdAt time.Time) Builder {
	builder.createdAt = createdAt
	return builder
}

func (builder *userBuilder) UpdatedAt(updatedAt time.Time) Builder {
	builder.updatedAt = updatedAt
	return builder
}

func (builder *userBuilder) UserStatus(status Status) Builder {
	builder.status = status
	return builder
}

func (builder *userBuilder) build() (*User, error) {
	var check = checkRequiredFieldsBuilder(builder)
	if check != nil {
		return nil, ErrCheckRequiredFieldsBuilder
	}

	return &User{
		id:           builder.id,
		name:         builder.name,
		telegramName: builder.telegramName,
		age:          builder.age,
		gender:       builder.gender,
		book:         builder.book,
		createdAt:    builder.createdAt,
		updatedAt:    builder.updatedAt,
		status:       builder.status,
	}, nil
}

func checkRequiredFieldsBuilder(builder *userBuilder) error {
	if len(builder.id.Value().String()) == 0 {
		return ErrEmptyUserIdBuilder
	}
	if len(builder.name.Value()) == 0 {
		return ErrEmptyNameBuilder
	}
	if len(builder.telegramName) == 0 {
		return ErrEmptyTelegramNameBuilder
	}
	if len(builder.book.BookID().Value().String()) == 0 {
		return ErrEmptyBookBuilder
	}
	if builder.createdAt.IsZero() {
		return ErrIsZeroTimeCreatedAtBuilder
	}
	if builder.updatedAt.IsZero() {
		return ErrIsZeroTimeUpdatedAtBuilder
	}

	return nil
}
