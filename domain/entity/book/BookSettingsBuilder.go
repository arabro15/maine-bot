package book

import (
	"errors"
	"time"
)

// Errors
var (
	ErrCheckRequiredFieldsSettingsBuilder = errors.New("check required fields in Builder is Error")
	ErrEmptyBookSettingsIDSettingsBuilder = errors.New("empty BookSettingsID in SettingsBuilder")
	ErrEmptyBookIDSettingsBuilder         = errors.New("empty BookID in SettingsBuilder")
	ErrEmptyPageInDaySettingsBuilder      = errors.New("empty PageInDay in SettingsBuilder")
	ErrIsZeroTimeUpdatedAtSettingsBuilder = errors.New("is zero time UpdatedAt in SettingsBuilder")
)

type SettingsBuilder struct {
	bookSettingsID SettingsID
	bookID         ID
	currentPage    int
	pageInDay      int
	updatedAt      time.Time
	bookStatus     Status
}

func NewSettingsBuilder() *SettingsBuilder {
	return &SettingsBuilder{}
}

func (sb *SettingsBuilder) SettingsID(id SettingsID) *SettingsBuilder {
	sb.bookSettingsID = id
	return sb
}

func (sb *SettingsBuilder) BookID(id ID) *SettingsBuilder {
	sb.bookID = id
	return sb
}

func (sb *SettingsBuilder) CurrentPage(currentPage int) *SettingsBuilder {
	sb.currentPage = currentPage
	return sb
}

func (sb *SettingsBuilder) PageInDay(pageInDay int) *SettingsBuilder {
	sb.pageInDay = pageInDay
	return sb
}

func (sb *SettingsBuilder) UpdateAt(updatedAt time.Time) *SettingsBuilder {
	sb.updatedAt = updatedAt
	return sb
}

func (sb *SettingsBuilder) BookStatus(status Status) *SettingsBuilder {
	sb.bookStatus = status
	return sb
}

func (sb *SettingsBuilder) Build() (*Settings, error) {
	var check = checkRequiredFieldsSettingsBuilder(sb)
	if check != nil {
		return nil, ErrCheckRequiredFieldsSettingsBuilder
	}
	return &Settings{
		bookSettingsID: sb.bookSettingsID,
		bookID:         sb.bookID,
		currentPage:    sb.currentPage,
		pageInDay:      sb.pageInDay,
		updatedAt:      sb.updatedAt,
		bookStatus:     sb.bookStatus,
	}, nil
}

func checkRequiredFieldsSettingsBuilder(sb *SettingsBuilder) error {
	if len(sb.bookSettingsID.Value().String()) == 0 {
		return ErrEmptyBookSettingsIDSettingsBuilder
	}
	if len(sb.bookID.Value().String()) == 0 {
		return ErrEmptyBookIDSettingsBuilder
	}
	if sb.pageInDay == 0 {
		return ErrEmptyPageInDaySettingsBuilder
	}
	if sb.updatedAt.IsZero() {
		return ErrIsZeroTimeUpdatedAtSettingsBuilder
	}

	return nil
}
