package entities

type Table struct {
	Id     int     `json:"id"`
	Places []Place `json:"places"`
}
