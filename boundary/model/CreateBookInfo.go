package model

type CreateBookInfo struct {
	UserID       string
	Name         string
	FilePath     string
	TextFilePath string
	BookSettings CreateBookSettingsInfo
	LoadedAt     string
}
