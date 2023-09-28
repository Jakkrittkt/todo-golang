package payload

import (
	"time"

	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/model"
)

type UpdateTodo struct {
	Title       string `binding:"required,max=100" validate:"required,max=100"`
	Description string
	Date        time.Time        `binding:"required"`
	Image       string           `validate:"imageBase64"`
	Status      model.TodoStatus `binding:"required,oneof=IN_PROGRESS COMPLETED" validate:"required,oneof=IN_PROGRESS COMPLETED"`
}
