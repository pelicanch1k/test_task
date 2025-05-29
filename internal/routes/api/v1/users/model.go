package users

type (
	CreateUser struct {
		Name       string `json:"name" validate:"required,max=40"`
		Surname    string `json:"surname" validate:"required,max=40"`
		Patronymic string `json:"patronymic" validate:"required,max=40"`
	}

	GetUsers struct {
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Age         int    `json:"age"`
		Age_2       int    `json:"age_2"`
		Gender      string `json:"gender"`
		Nationality string `json:"nationality"`
		Limit       int    `json:"limit"`
		Offset      int    `json:"offset"`
	}

	DeleteUser struct {
		Id int32 `json:"id validate:"required`
	}
)
