package model

type CreateUserParams struct {
	Name string
	Surname string
	Patronymic string
}

type GetUsersParams struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Age         int    `json:"age"`
	Age_2       int    `json:"age_2"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

