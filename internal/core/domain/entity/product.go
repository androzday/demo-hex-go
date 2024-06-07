package entity

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}
