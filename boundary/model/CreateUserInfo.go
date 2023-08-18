package model

type CreateUserInfo struct {
	name         string
	telegramName string
	age          string
	gender       string
	createdAt    string
	updatedAt    string
	status       string
}

func NewCreateUserInfo(name string, telegramName string, age string, gender string, createdAt string, updatedAt string, status string) *CreateUserInfo {
	return &CreateUserInfo{name: name, telegramName: telegramName, age: age, gender: gender, createdAt: createdAt, updatedAt: updatedAt, status: status}
}

func (c *CreateUserInfo) Name() string {
	return c.name
}

func (c *CreateUserInfo) SetName(name string) {
	c.name = name
}

func (c *CreateUserInfo) TelegramName() string {
	return c.telegramName
}

func (c *CreateUserInfo) SetTelegramName(telegramName string) {
	c.telegramName = telegramName
}

func (c *CreateUserInfo) Age() string {
	return c.age
}

func (c *CreateUserInfo) SetAge(age string) {
	c.age = age
}

func (c *CreateUserInfo) Gender() string {
	return c.gender
}

func (c *CreateUserInfo) SetGender(gender string) {
	c.gender = gender
}

func (c *CreateUserInfo) CreatedAt() string {
	return c.createdAt
}

func (c *CreateUserInfo) SetCreatedAt(createdAt string) {
	c.createdAt = createdAt
}

func (c *CreateUserInfo) UpdatedAt() string {
	return c.updatedAt
}

func (c *CreateUserInfo) SetUpdatedAt(updatedAt string) {
	c.updatedAt = updatedAt
}

func (c *CreateUserInfo) Status() string {
	return c.status
}

func (c *CreateUserInfo) SetStatus(status string) {
	c.status = status
}
