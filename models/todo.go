package model

type Todo struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    string `json:"userId"`
}

type PostTodo struct {
	Title     string `json:"title"`
	UserId    string `json:"userId"`
	Completed bool   `json:"completed"`
}

type TodoUri struct {
	ID uint `uri:"id" binding:"required,number"`
}
