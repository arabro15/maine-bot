package user

import (
	"errors"
	"maineBot/domain/entity"
	"time"
)

// Errors
var (
	ErrCheckRequiredFieldsBuilder = errors.New("check required fields in Builder is Error")
	ErrEmptyUserIdBuilder         = errors.New("empty UserID in Builder")
	ErrEmptyNameBuilder           = errors.New("empty Name in Builder")
	ErrEmptyTelegramNameBuilder   = errors.New("empty TelegramName in Builder")
	ErrIsZeroTimeCreatedAtBuilder = errors.New("time CreatedAt is Zero in Builder")
	ErrIsZeroTimeUpdatedAtBuilder = errors.New("time UpdatedAt is Zero in Builder")
)

type Errors []error

func (e *Errors) Error() string {
	return ""
}

type Builder struct {
	id           ID
	name         entity.Name
	telegramName string
	age          Age
	gender       Gender
	createdAt    time.Time
	updatedAt    time.Time
	status       Status
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) UserID(id ID) *Builder {
	b.id = id
	return b
}

func (b *Builder) Name(name entity.Name) *Builder {
	b.name = name
	return b
}

func (b *Builder) TelegramName(telegramName string) *Builder {
	b.telegramName = telegramName
	return b
}

func (b *Builder) Age(age Age) *Builder {
	b.age = age
	return b
}

func (b *Builder) Gender(gender Gender) *Builder {
	b.gender = gender
	return b
}

func (b *Builder) CreatedAt(createdAt time.Time) *Builder {
	b.createdAt = createdAt
	return b
}

func (b *Builder) UpdatedAt(updatedAt time.Time) *Builder {
	b.updatedAt = updatedAt
	return b
}

func (b *Builder) UserStatus(status Status) *Builder {
	b.status = status
	return b
}

func (b *Builder) Build() (User, error) {
	var check = checkRequiredFieldsBuilder(b)
	if check != nil {
		return User{}, ErrCheckRequiredFieldsBuilder
	}

	return User{
		id:           b.id,
		name:         b.name,
		telegramName: b.telegramName,
		age:          b.age,
		gender:       b.gender,
		createdAt:    b.createdAt,
		updatedAt:    b.updatedAt,
		status:       b.status,
	}, nil
}

func checkRequiredFieldsBuilder(b *Builder) error {
	errs := Errors{}

	if len(b.id.Value().String()) == 0 {
		errs = append(errs, ErrEmptyUserIdBuilder)
	}
	if len(b.name.Value()) == 0 {
		return ErrEmptyNameBuilder
	}
	if len(b.telegramName) == 0 {
		return ErrEmptyTelegramNameBuilder
	}
	if b.createdAt.IsZero() {
		return ErrIsZeroTimeCreatedAtBuilder
	}
	if b.updatedAt.IsZero() {
		return ErrIsZeroTimeUpdatedAtBuilder
	}

	return nil
}
