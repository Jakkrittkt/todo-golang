package payload

import "github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/model"

type CreateTodo struct {
	Title       string `binding:"required,max=100" validate:"required,max=100"`
	Description string
	Image       string           `validate:"imageBase64"`
	Status      model.TodoStatus `binding:"required,oneof=IN_PROGRESS COMPLETED" validate:"required,oneof=IN_PROGRESS COMPLETED"`
}
