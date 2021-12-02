package data

import "A-Golang-MiniProject/features/spoonacular"

type Response struct {
	Message    interface{} `json:"message"`
	HealtyFood []struct {
		Id        string `json:"id"`
		Title     string `json:"title"`
		Image     string `json:"image"`
		ImageType string `json:"imageType"`
		Calories  int    `json:"calories"`
		Protein   string `json:"protein"`
		Fat       string `json:"fat"`
		Carbs     string `json:"carbs"`
	} `json:"healtyFood"`
}

type HealtyFood []struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	ImageType string `json:"imageType"`
	Calories  int    `json:"calories"`
	Protein   string `json:"protein"`
	Fat       string `json:"fat"`
	Carbs     string `json:"carbs"`
}

func (resp *Response) ToDomain() spoonacular.Core {
	return spoonacular.Core{
		HealtyFood: resp.HealtyFood,
	}
}

func ToListDomain(data []Response) []spoonacular.Core {
	result := []spoonacular.Core{}
	for _, domain := range data {
		result = append(result, domain.ToDomain())
	}
	return result
}
