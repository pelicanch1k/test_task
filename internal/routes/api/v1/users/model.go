package users

type (
	CreateUser struct {
		Name string `json:"name" validate:"required,max=40"`
		Surname string `json:"surname" validate:"required,max=40"`
		Patronymic string `json:"patronymic" validate:"required,max=40"`
	}
)