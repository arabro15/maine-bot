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

type SettingsBuilder interface {
	SettingsID(id SettingsID) SettingsBuilder
	BookID(id ID) SettingsBuilder
	CurrentPage(currentPage int) SettingsBuilder
	PageInDay(pageInDay int) SettingsBuilder
	UpdateAt(updatedAt time.Time) SettingsBuilder
	BookStatus(status Status) SettingsBuilder
	build() (*Settings, error)
}
type settingsBuilder struct {
	bookSettingsID SettingsID
	bookID         ID
	currentPage    int
	pageInDay      int
	updatedAt      time.Time
	bookStatus     Status
}

func (builder *settingsBuilder) SettingsID(id SettingsID) SettingsBuilder {
	builder.bookSettingsID = id
	return builder
}

func (builder *settingsBuilder) BookID(id ID) SettingsBuilder {
	builder.bookID = id
	return builder
}

func (builder *settingsBuilder) CurrentPage(currentPage int) SettingsBuilder {
	builder.currentPage = currentPage
	return builder
}

func (builder *settingsBuilder) PageInDay(pageInDay int) SettingsBuilder {
	builder.pageInDay = pageInDay
	return builder
}

func (builder *settingsBuilder) UpdateAt(updatedAt time.Time) SettingsBuilder {
	builder.updatedAt = updatedAt
	return builder
}

func (builder *settingsBuilder) BookStatus(status Status) SettingsBuilder {
	builder.bookStatus = status
	return builder
}

func (builder *settingsBuilder) build() (*Settings, error) {
	var check = checkRequiredFieldsSettingsBuilder(builder)
	if check != nil {
		return nil, ErrCheckRequiredFieldsSettingsBuilder
	}
	return &Settings{
		bookSettingsID: builder.bookSettingsID,
		bookID:         builder.bookID,
		currentPage:    builder.currentPage,
		pageInDay:      builder.pageInDay,
		updatedAt:      builder.updatedAt,
		bookStatus:     builder.bookStatus,
	}, nil
}

func checkRequiredFieldsSettingsBuilder(builder *settingsBuilder) error {
	if len(builder.bookSettingsID.Value().String()) == 0 {
		return ErrEmptyBookSettingsIDSettingsBuilder
	}
	if len(builder.bookID.Value().String()) == 0 {
		return ErrEmptyBookIDSettingsBuilder
	}
	if builder.pageInDay == 0 {
		return ErrEmptyPageInDaySettingsBuilder
	}
	if builder.updatedAt.IsZero() {
		return ErrIsZeroTimeUpdatedAtSettingsBuilder
	}

	return nil
}
