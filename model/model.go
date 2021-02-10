package model

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Todos []Todos `json:"todos"`
}
type Todos struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}
