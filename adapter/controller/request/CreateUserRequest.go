package request

type CreateUserRequest struct {
	name         string
	telegramName string
	age          string
	gender       string
	createdAt    string
	updatedAt    string
	status       string
}

func NewCreateUserRequest(
	name string,
	telegramName string,
	age string, gender string,
	createdAt string,
	updatedAt string,
	status string,
) *CreateUserRequest {
	return &CreateUserRequest{
		name:         name,
		telegramName: telegramName,
		age:          age,
		gender:       gender,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		status:       status,
	}
}

func (c *CreateUserRequest) Name() string {
	return c.name
}

func (c *CreateUserRequest) SetName(name string) {
	c.name = name
}

func (c *CreateUserRequest) TelegramName() string {
	return c.telegramName
}

func (c *CreateUserRequest) SetTelegramName(telegramName string) {
	c.telegramName = telegramName
}

func (c *CreateUserRequest) Age() string {
	return c.age
}

func (c *CreateUserRequest) SetAge(age string) {
	c.age = age
}

func (c *CreateUserRequest) Gender() string {
	return c.gender
}

func (c *CreateUserRequest) SetGender(gender string) {
	c.gender = gender
}

func (c *CreateUserRequest) CreatedAt() string {
	return c.createdAt
}

func (c *CreateUserRequest) SetCreatedAt(createdAt string) {
	c.createdAt = createdAt
}

func (c *CreateUserRequest) UpdatedAt() string {
	return c.updatedAt
}

func (c *CreateUserRequest) SetUpdatedAt(updatedAt string) {
	c.updatedAt = updatedAt
}

func (c *CreateUserRequest) Status() string {
	return c.status
}

func (c *CreateUserRequest) SetStatus(status string) {
	c.status = status
}
