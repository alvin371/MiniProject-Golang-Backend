package data

import (
	"A-Golang-MiniProject/features/spoonacular"
	"encoding/json"
	"net/http"
	"strconv"
)

type Api struct {
	httpClient http.Client
}

func NewApi() spoonacular.Data {
	return &Api{
		httpClient: http.Client{},
	}
}

func (api *Api) GetByNutrients(random bool, minCarbs int32, number int16) (spoonacular.Core, error) {

	client := &http.Client{}
	apiKey := "fc2d34fd7a7445d38379cfceb5cf396a"
	randomize := strconv.FormatBool(random)
	carb := strconv.Itoa(int(minCarbs))
	total := strconv.Itoa(int(number))
	url := "https://api.spoonacular.com/recipes/findByNutrients/?random=" + randomize + "&minCarbs=" + carb + "&number=" + total + "&apiKey=" + apiKey
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		return spoonacular.Core{}, err
	}

	data := Response{}

	err = json.NewDecoder(resp.Body).Decode(&data)

	return data.ToDomain(), err
}
