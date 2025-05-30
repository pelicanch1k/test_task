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

	UpdateUser struct {
		Name        *string  `json:"name,omitempty"`
		Surname     *string  `json:"surname, omitempty"`
		Patronymic  *string `json:"patronymic,omitempty"`
		Age         *int32  `json:"age,omitempty" validate:"omitempty,min=0,max=150"`
		Gender      *string `json:"gender,omitempty" validate:"omitempty,oneof=male female other"`
		Nationality *string `json:"nationality,omitempty" validate:"omitempty,min=2,max=2"`
	}
)
