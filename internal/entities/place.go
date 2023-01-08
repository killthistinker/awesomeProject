package entities

type Place struct {
	Id     int    `json:"id"`
	Number int    `json:"number"`
	Dishes []Dish `json:"dishes"`
}
