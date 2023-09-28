package repo

import (
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/model"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/payload"
)

type TodoRepository interface {
	New(*model.Todo) error
	GetById(*model.Todo) error
	List(*[]model.Todo, *payload.FilterTodo) error
	Update(*model.Todo) error
}
