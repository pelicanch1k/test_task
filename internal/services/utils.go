package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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