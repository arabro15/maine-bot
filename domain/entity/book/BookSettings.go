package book

import (
	"errors"
	"time"
)

var (
	ErrEmptyBookSettingsID             = errors.New("empty BookSettingsID")
	ErrEmptyBookSettingsBookID         = errors.New("empty BookID in BookSettings")
	ErrEmptyPageInDay                  = errors.New("empty PageInDay")
	ErrIsZeroBookSettingsUpdateAt      = errors.New("is zero time UpdateAt in BookSettings")
	ErrBookSettingsCheckRequiredFields = errors.New("check required fields in BookSettings is error")
)

type Settings struct {
	bookSettingsID SettingsID
	bookID         ID
	currentPage    int
	pageInDay      int
	updatedAt      time.Time
	bookStatus     Status
}

func (bs Settings) NewBookSettings(
	bookSettingsID SettingsID,
	bookID ID,
	pageInDay int,
	updatedAt time.Time,
	bookStatus Status,
) (*Settings, error) {
	var check = bs.checkRequiredFields(
		&bookSettingsID,
		&bookID,
		&pageInDay,
		&updatedAt,
	)
	if check != nil {
		return nil, ErrBookSettingsCheckRequiredFields
	}

	return &Settings{
		bookSettingsID: bookSettingsID,
		bookID:         bookID,
		pageInDay:      pageInDay,
		updatedAt:      updatedAt,
		bookStatus:     bookStatus,
	}, nil
}

func (bs Settings) checkRequiredFields(
	bookSettingsID *SettingsID,
	bookID *ID,
	pageInDay *int,
	updatedAt *time.Time,
) error {
	if len(bookSettingsID.Value().String()) == 0 {
		return ErrEmptyBookSettingsID
	}
	if len(bookID.Value().String()) == 0 {
		return ErrEmptyBookSettingsBookID
	}
	if *pageInDay == 0 {
		return ErrEmptyPageInDay
	}
	if updatedAt.IsZero() {
		return ErrIsZeroBookSettingsUpdateAt
	}

	return nil
}
