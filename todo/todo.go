package todo

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/model"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/payload"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/repo"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func CustomValidateImageBase64(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	fmt.Println()

	if !strings.HasPrefix(value, "data:image") {
		return false
	}

	dataParts := strings.Split(value, ",")
	if len(dataParts) != 2 {
		return false
	}

	_, err := base64.StdEncoding.DecodeString(dataParts[1])
	return err == nil
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	Param(string) string
	Query(string) string
}

type TodoHandler struct {
	todoRepo repo.TodoRepository
}

func NewTodoHandler(todoRepo repo.TodoRepository) *TodoHandler {
	return &TodoHandler{todoRepo: todoRepo}
}

func (t *TodoHandler) Create(c Context) {
	var payload payload.CreateTodo

	if err := c.Bind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	todo := model.Todo{
		ID:          uuid.New(),
		Title:       payload.Title,
		Image:       payload.Image,
		Description: payload.Description,
		Status:      payload.Status,
		Date:        time.Now(),
	}

	if err := t.todoRepo.New(&todo); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":          todo.ID,
		"title":       todo.Title,
		"description": todo.Description,
		"date":        todo.Date,
		"image":       todo.Image,
		"status":      todo.Status,
	})
}

func (t *TodoHandler) FindById(c Context) {

	var todo model.Todo

	id, errParse := uuid.Parse(c.Param("id"))
	if errParse != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": errParse.Error(),
		})
		return
	}

	todo.ID = id
	err := t.todoRepo.GetById(&todo)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (t *TodoHandler) FindAll(c Context) {
	var todos []model.Todo

	filter := payload.FilterTodo{
		SortBy:  c.Query("sort"),
		OrderBy: c.Query("order"),
		Search:  c.Query("search"),
	}

	err := t.todoRepo.List(&todos, &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) UpdateById(c Context) {
	var todo model.Todo
	var payload payload.UpdateTodo

	id, errParse := uuid.Parse(c.Param("id"))
	if errParse != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": errParse.Error(),
		})
		return
	}

	if err := c.Bind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	todo.ID = id
	todo.Title = payload.Title
	todo.Description = payload.Description
	todo.Date = payload.Date
	todo.Image = payload.Image
	todo.Status = payload.Status
	err := t.todoRepo.Update(&todo)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}
