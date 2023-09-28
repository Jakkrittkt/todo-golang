package db

import (
	"errors"
	"os"
	"time"

	"encoding/json"

	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/model"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/payload"
	"github.com/google/uuid"
	"github.com/thedevsaddam/gojsonq"
)

type DbJsonFile struct {
	fileName string
}

func NewDbJsonFile(fileName string) *DbJsonFile {
	file, err := os.Open(fileName)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}

		file.WriteString("[]")
	}
	defer file.Close()
	return &DbJsonFile{fileName: fileName}
}

func (db *DbJsonFile) New(todo *model.Todo) error {
	var todos []model.Todo

	oldData, err := os.ReadFile(db.fileName)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(oldData, &todos); err != nil {
		return err
	}

	todos = append(todos, *todo)
	newData, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(db.fileName, newData, 0644); err != nil {
		return err
	}

	return nil
}
func (db *DbJsonFile) GetById(todo *model.Todo) error {
	var todos []model.Todo
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &todos); err != nil {
		return err
	}

	found := false
	for _, value := range todos {
		if value.ID == todo.ID {
			todo.Title = value.Title
			todo.Description = value.Description
			todo.Date = value.Date
			todo.Image = value.Image
			todo.Status = value.Status
			found = true
			break
		}
	}

	if !found {
		return errors.New("record not found")
	}

	return nil
}
func (db *DbJsonFile) List(todos *[]model.Todo, filter *payload.FilterTodo) error {

	jq := gojsonq.New().File(db.fileName)

	if filter.Search != "" {
		jq = jq.Where("title", "contains", filter.Search).OrWhere("description", "contains", filter.Search)
	}
	if filter.SortBy != "" {
		jq = jq.SortBy(filter.SortBy, filter.OrderBy)
	}

	res := jq.Get()

	for _, value := range res.([]interface{}) {
		date, _ := time.Parse(time.RFC3339, value.(map[string]interface{})["date"].(string))
		*todos = append(*todos, model.Todo{
			ID:          uuid.MustParse(value.(map[string]interface{})["id"].(string)),
			Title:       value.(map[string]interface{})["title"].(string),
			Description: value.(map[string]interface{})["description"].(string),
			Date:        date,
			Status:      model.TodoStatus(value.(map[string]interface{})["date"].(string)),
		})

	}
	return nil
}
func (db *DbJsonFile) Update(todo *model.Todo) error {
	var todos []model.Todo
	oldData, err := os.ReadFile(db.fileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(oldData, &todos); err != nil {
		return err
	}

	found := false
	for index, value := range todos {
		if todo.ID == value.ID {
			todos[index].Title = todo.Title
			todos[index].Description = todo.Description
			todos[index].Date = todo.Date
			todos[index].Image = todo.Image
			todos[index].Status = todo.Status
			found = true
			break
		}
	}

	if !found {
		return errors.New("record not found")
	}

	newData, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(db.fileName, newData, 0644); err != nil {
		return err
	}

	return nil
}
