package book

import (
	"errors"
	"maineBot/domain/entity"
	"maineBot/domain/entity/user"
	"time"
)

// Errors
var (
	ErrEmptyBookID             = errors.New("empty BookID")
	ErrEmptyBookUserID         = errors.New("empty UserID in Book")
	ErrEmptyBookName           = errors.New("empty Name in Book")
	ErrEmptyFilePath           = errors.New("empty filepath in Book")
	ErrEmptyBookSettings       = errors.New("empty BookSettings")
	ErrIsZeroTimeLoadedAt      = errors.New("is zero time loadedAt in Book")
	ErrBookCheckRequiredFields = errors.New("func checkRequiredFields is Error")
)

type Book struct {
	bookID       ID
	userID       user.ID
	name         entity.Name
	filePath     string
	textFilePath string
	bookSettings Settings
	loadedAt     time.Time
}

func (b Book) NewBook(
	bookID ID,
	userID user.ID,
	name entity.Name,
	filePath string,
	bookSettings Settings,
	loadedAt time.Time,
) (*Book, error) {
	var check = b.checkRequiredFields(&bookID, &userID, &name, &filePath, &bookSettings, &loadedAt)
	if check != nil {
		return nil, ErrBookCheckRequiredFields
	}

	return &Book{
		bookID:       bookID,
		userID:       userID,
		name:         name,
		filePath:     filePath,
		bookSettings: bookSettings,
		loadedAt:     loadedAt,
	}, nil
}

func (b Book) BookID() *ID {
	return &b.bookID
}

func (b Book) UserID() user.ID {
	return b.userID
}

func (b Book) Name() *entity.Name {
	return &b.name
}

func (b Book) FilePath() string {
	return b.filePath
}

func (b Book) TextFilePath() string {
	return b.textFilePath
}

func (b Book) BookSettings() Settings {
	return b.bookSettings
}

func (b Book) LoadedAt() time.Time {
	return b.loadedAt
}

func (b Book) checkRequiredFields(
	bookID *ID,
	userID *user.ID,
	name *entity.Name,
	filePath *string,
	bookSettings *Settings,
	loadedAt *time.Time,
) error {
	if len(bookID.Value().String()) == 0 {
		return ErrEmptyBookID
	}
	if len(userID.Value().String()) == 0 {
		return ErrEmptyBookUserID
	}
	if len(name.Value()) == 0 {
		return ErrEmptyBookName
	}
	if len(*filePath) == 0 {
		return ErrEmptyFilePath
	}
	if len(bookSettings.bookSettingsID.Value().String()) == 0 {
		return ErrEmptyBookSettings
	}
	if loadedAt.IsZero() {
		return ErrIsZeroTimeLoadedAt
	}

	return nil
}
