package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"test_task/internal/model"
	"test_task/internal/repository/gen"
)

func GetRequest(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при отправке запроса: %s\n", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %s\n", err)
		return nil
	}

	// Парсим в map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Ошибка при парсинге JSON: %s\n", err)
		return nil
	}

	return result
}

func ConvertToGetUsersParams(input *model.GetUsersParams) *gen.GetUsersParams {
	var name interface{}
	if input.Name != "" {
		name = input.Name
	}

	var surname interface{}
	if input.Surname != "" {
		surname = input.Surname
	}

	var minAge interface{}
	if input.Age > 0 {
		minAge = input.Age
	}

	var maxAge interface{}
	if input.Age_2 > 0 {
		maxAge = input.Age_2
	}

	var gender interface{}
	if input.Gender != "" {
		gender = input.Gender
	}

	var nationality interface{}
	if input.Nationality != "" {
		nationality = input.Nationality
	}

	var limit, offset int32
	if input.Limit <= 0 {
		limit = 10
	}

	if input.Offset <= 0 {
		offset = 1
	}

	return &gen.GetUsersParams{
		Name:        name,        // Column1
		Surname:     surname,     // Column2
		MinAge:      minAge,      // Column3 (age >=)
		MaxAge:      maxAge,      // Column4 (age <=)
		Gender:      gender,      // Column5
		Nationality: nationality, // Column6
		Limit:       limit,
		Offset:      offset,
	}
}

func interfaceToIntViaString(val interface{}) (int, error) {
    str := fmt.Sprintf("%v", val)
    return strconv.Atoi(str)
}