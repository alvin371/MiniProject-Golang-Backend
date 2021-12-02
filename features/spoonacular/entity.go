package spoonacular

type Core struct {
	HealtyFood interface{}
}

type Data interface {
	GetByNutrients(random bool, minCarbs int32, number int16) (Core, error)
}
