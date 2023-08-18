package model

import "github.com/google/uuid"

type BookSettingsDbModel struct {
	bookSettingsID uuid.UUID
	bookID         uuid.UUID
	currentPage    int
	pageInDay      int
	updatedAt      string
	bookStatus     string
}

func NewBookSettingsDbModel(bookSettingsID uuid.UUID,
	bookID uuid.UUID,
	currentPage int,
	pageInDay int,
	updatedAt string,
	bookStatus string,
) *BookSettingsDbModel {
	return &BookSettingsDbModel{
		bookSettingsID: bookSettingsID,
		bookID:         bookID,
		currentPage:    currentPage,
		pageInDay:      pageInDay,
		updatedAt:      updatedAt,
		bookStatus:     bookStatus,
	}
}
