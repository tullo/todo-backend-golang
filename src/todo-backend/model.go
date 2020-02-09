package main

// Todo is a structure that models todo information 
type Todo struct {
	ID        int    `json:"-"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	URL       string `json:"url"`
}
