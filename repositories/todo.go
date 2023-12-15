package repositories

import (
	"database/sql"
	model "go_app1/models"
	"log"
)

type TodoRepository struct {
	Db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepositoryInterface {
	return &TodoRepository{Db: db}
}

// DeleteTodo implements TodoRepositoryInterface
func (m *TodoRepository) DeleteTodo(id uint) bool {

	count := 0

	err := m.Db.QueryRow("SELECT COUNT(*) FROM todos WHERE id = $1", id).Scan(&count)
	if err != nil {
		return false
	}

	if count == 0 {
		return false
	}

	_, err = m.Db.Exec("DELETE FROM todos WHERE id = $1", id)
	// _, err := m.Db.Exec("DELETE FROM todo WHERE id = $1", id)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GetAllTodo implements TodoRepositoryInterface
func (m *TodoRepository) GetAllTodo() []model.Todo {
	query, err := m.Db.Query("SELECT * FROM todos")
	if err != nil {
		log.Println(err)
		return nil
	}

	var todos []model.Todo
	if query != nil {
		for query.Next() {
			var (
				id        uint
				title     string
				completed bool
				userId    int
			)
			err := query.Scan(&id, &title, &completed, &userId)
			if err != nil {
				log.Println(err)
			}
			todo := model.Todo{Id: id, Title: title, UserId: userId}
			todos = append(todos, todo)
		}
	}
	return todos
}

// GetOneTodo implements TodoRepositoryInterface
func (m *TodoRepository) GetOneTodo(id uint) model.Todo {
	query, err := m.Db.Query("SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return model.Todo{}
	}
	var todo model.Todo
	if query != nil {
		for query.Next() {
			var (
				id        uint
				title     string
				completed bool
				userId    int
			)
			err := query.Scan(&id, &title, &completed, &userId)
			if err != nil {
				log.Println(err)
			}
			todo = model.Todo{Id: id, Title: title, Completed: completed, UserId: userId}
		}
	}
	return todo
}

// InsertTodo implements TodoRepositoryInterface
func (m *TodoRepository) InsertTodo(post model.PostTodo) bool {
	stmt, err := m.Db.Prepare("INSERT INTO todos(title, completed, userId) VALUES ($1,$2,$3)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Title, post.Completed, post.UserId)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// UpdateTodo implements TodoRepositoryInterface
func (m *TodoRepository) UpdateTodo(id uint, post model.PostTodo) model.Todo {
	_, err := m.Db.Exec("UPDATE todos SET title = $1, completed = $2, userId = $3 WHERE id = $4", post.Title, post.Completed, post.UserId, id)
	if err != nil {
		log.Println(err)
		return model.Todo{}
	}
	return m.GetOneTodo(id)
}
