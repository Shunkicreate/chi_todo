package entities

type Todo struct {
	ID        int `json:"id"`
	Title     string `json:"title"`
	Discription string `json:"discription"`
	Completed bool `json:"completed"`
	Link 	string `json:"link"`
}
