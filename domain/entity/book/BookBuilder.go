package book

import (
	"errors"
	"maineBot/domain/entity"
	"maineBot/domain/entity/user"
	"time"
)

// Errors
var (
	ErrEmptyBookIdBuilder         = errors.New("empty BookID in Builder")
	ErrEmptyBookUserIdBuilder     = errors.New("empty Book-UserID in Builder")
	ErrEmptyBookNameBuilder       = errors.New("empty Book-Name in Builder")
	ErrEmptyFilePathBuilder       = errors.New("empty filePath in Builder")
	ErrEmptyTextFilePathBuilder   = errors.New("empty TextFilePath in Builder")
	ErrEmptyBookSettingsBuilder   = errors.New("empty BookSettings in Builder")
	ErrIsZeroTimeLoadedAtBuilder  = errors.New("is zero time LoadedAt in Builder")
	ErrCheckRequiredFieldsBuilder = errors.New("check required fields in Builder is Error")
)

type Builder struct {
	bookID       ID
	userID       user.ID
	name         entity.Name
	filePath     string
	textFilePath string
	bookSettings Settings
	loadedAt     time.Time
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) BookID(id ID) *Builder {
	b.bookID = id
	return b
}

func (b *Builder) UserID(userID user.ID) *Builder {
	b.userID = userID
	return b
}

func (b *Builder) Name(name entity.Name) *Builder {
	b.name = name
	return b
}

func (b *Builder) FilePath(filePath string) *Builder {
	b.filePath = filePath
	return b
}

func (b *Builder) TextFilePath(textFilePath string) *Builder {
	b.textFilePath = textFilePath
	return b
}

func (b *Builder) BookSettings(bookSettings Settings) *Builder {
	b.bookSettings = bookSettings
	return b
}

func (b *Builder) LoadedAt(loadedAt time.Time) *Builder {
	b.loadedAt = loadedAt
	return b
}

func (b *Builder) Build() (*Book, error) {
	var check = checkRequiredFieldsBuilder(b)
	if check != nil {
		return nil, ErrCheckRequiredFieldsBuilder
	}
	return &Book{
		bookID:       b.bookID,
		userID:       b.userID,
		name:         b.name,
		filePath:     b.filePath,
		textFilePath: b.textFilePath,
		bookSettings: b.bookSettings,
		loadedAt:     b.loadedAt,
	}, nil
}

func checkRequiredFieldsBuilder(b *Builder) error {
	if len(b.bookID.Value().String()) == 0 {
		return ErrEmptyBookIdBuilder
	}
	if len(b.userID.Value().String()) == 0 {
		return ErrEmptyBookUserIdBuilder
	}
	if len(b.name.Value()) == 0 {
		return ErrEmptyBookNameBuilder
	}
	if len(b.filePath) == 0 {
		return ErrEmptyFilePathBuilder
	}
	if len(b.textFilePath) == 0 {
		return ErrEmptyTextFilePathBuilder
	}
	if len(b.bookSettings.bookSettingsID.Value().String()) == 0 {
		return ErrEmptyBookSettingsBuilder
	}
	if b.loadedAt.IsZero() {
		return ErrIsZeroTimeLoadedAtBuilder
	}

	return nil
}
