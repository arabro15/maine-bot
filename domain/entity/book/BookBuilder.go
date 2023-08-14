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

type Builder interface {
	BookID(id ID) Builder
	UserID(userID user.ID) Builder
	Name(name entity.Name) Builder
	FilePath(filePath string) Builder
	TextFilePath(textFilePath string) Builder
	BookSettings(bookSettings Settings) Builder
	LoadedAt(loadedAt time.Time) Builder
	build() (*Book, error)
}

type builder struct {
	bookID       ID
	userID       user.ID
	name         entity.Name
	filePath     string
	textFilePath string
	bookSettings Settings
	loadedAt     time.Time
}

func (builder *builder) BookID(id ID) Builder {
	builder.bookID = id
	return builder
}

func (builder *builder) UserID(userID user.ID) Builder {
	builder.userID = userID
	return builder
}

func (builder *builder) Name(name entity.Name) Builder {
	builder.name = name
	return builder
}

func (builder *builder) FilePath(filePath string) Builder {
	builder.filePath = filePath
	return builder
}

func (builder *builder) TextFilePath(textFilePath string) Builder {
	builder.textFilePath = textFilePath
	return builder
}

func (builder *builder) BookSettings(bookSettings Settings) Builder {
	builder.bookSettings = bookSettings
	return builder
}

func (builder *builder) LoadedAt(loadedAt time.Time) Builder {
	builder.loadedAt = loadedAt
	return builder
}

func (builder *builder) build() (*Book, error) {
	var check = checkRequiredFieldsBuilder(builder)
	if check != nil {
		return nil, ErrCheckRequiredFieldsBuilder
	}
	return &Book{
		bookID:       builder.bookID,
		userID:       builder.userID,
		name:         builder.name,
		filePath:     builder.filePath,
		textFilePath: builder.textFilePath,
		bookSettings: builder.bookSettings,
		loadedAt:     builder.loadedAt,
	}, nil
}

func checkRequiredFieldsBuilder(builder *builder) error {
	if len(builder.bookID.Value().String()) == 0 {
		return ErrEmptyBookIdBuilder
	}
	if len(builder.userID.Value().String()) == 0 {
		return ErrEmptyBookUserIdBuilder
	}
	if len(builder.name.Value()) == 0 {
		return ErrEmptyBookNameBuilder
	}
	if len(builder.filePath) == 0 {
		return ErrEmptyFilePathBuilder
	}
	if len(builder.textFilePath) == 0 {
		return ErrEmptyTextFilePathBuilder
	}
	if len(builder.bookSettings.bookSettingsID.Value().String()) == 0 {
		return ErrEmptyBookSettingsBuilder
	}
	if builder.loadedAt.IsZero() {
		return ErrIsZeroTimeLoadedAtBuilder
	}

	return nil
}
