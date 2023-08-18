package model

import (
	"github.com/google/uuid"
)

type UserDbModel struct {
	id           uuid.UUID
	name         string
	telegramName string
	age          string
	gender       string
	createdAt    string
	updatedAt    string
	status       string
}

func NewUserDbModel(
	id uuid.UUID,
	name string,
	telegramName string,
	age string,
	gender string,
	createdAt string,
	updatedAt string,
	status string,
) *UserDbModel {
	return &UserDbModel{
		id:           id,
		name:         name,
		telegramName: telegramName,
		age:          age,
		gender:       gender,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		status:       status,
	}
}

func (m *UserDbModel) GetId() uuid.UUID {
	return m.id
}

func (m *UserDbModel) SetId(id uuid.UUID) {
	m.id = id
}

func (m *UserDbModel) GetName() string {
	return m.name
}

func (m *UserDbModel) SetName(name string) {
	m.name = name
}

func (m *UserDbModel) GetTelegramName() string {
	return m.telegramName
}

func (m *UserDbModel) SetTelegramName(telegramName string) {
	m.telegramName = telegramName
}

func (m *UserDbModel) GetAge() string {
	return m.age
}

func (m *UserDbModel) SetAge(age string) {
	m.age = age
}

func (m *UserDbModel) GetGender() string {
	return m.gender
}

func (m *UserDbModel) SetGender(gender string) {
	m.gender = gender
}

func (m *UserDbModel) GetCreatedAt() string {
	return m.createdAt
}

func (m *UserDbModel) SetCreatedAt(createdAt string) {
	m.createdAt = createdAt
}

func (m *UserDbModel) GetUpdatedAt() string {
	return m.updatedAt
}

func (m *UserDbModel) SetUpdatedAt(updatedAt string) {
	m.updatedAt = updatedAt
}

func (m *UserDbModel) GetStatus() string {
	return m.status
}

func (m *UserDbModel) SetStatus(status string) {
	m.status = status
}
