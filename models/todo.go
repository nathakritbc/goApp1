package model

type Todo struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

type PostTodo struct {
	Title     string `json:"title"`
	UserId    int    `json:"userId"`
	Completed bool   `json:"completed"`
}

type TodoUri struct {
	ID uint `uri:"id" binding:"required,number"`
}
