package models

type Pizza struct {
	Id           int      `json:"id"`
	Nome         string   `json:"nome"`
	Ingredientes []string `json:"ingredientes"`
	Preco        float64  `json:"preco"`
}
