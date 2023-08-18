package model

import (
	"github.com/google/uuid"
)

type BookDbModel struct {
	bookID       uuid.UUID
	userID       uuid.UUID
	name         string
	filePath     string
	textFilePath string
	bookSettings BookSettingsDbModel
	loadedAt     string
}

func NewBookDbModel(
	bookID uuid.UUID,
	userID uuid.UUID,
	name string,
	filePath string,
	textFilePath string,
	bookSettings BookSettingsDbModel,
	loadedAt string,
) *BookDbModel {
	return &BookDbModel{
		bookID:       bookID,
		userID:       userID,
		name:         name,
		filePath:     filePath,
		textFilePath: textFilePath,
		bookSettings: bookSettings,
		loadedAt:     loadedAt,
	}
}
