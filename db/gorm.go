package db

import (
	"fmt"

	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/model"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/payload"
	"gorm.io/gorm"
)

type Gorm struct {
	db *gorm.DB
}

func NewGorm(db *gorm.DB) *Gorm {
	return &Gorm{db: db}
}

func (store *Gorm) New(todo *model.Todo) error {
	return store.db.Create(&todo).Error
}

func (store *Gorm) GetById(todo *model.Todo) error {
	return store.db.First(&todo).Error
}

func (store *Gorm) List(todo *[]model.Todo, filter *payload.FilterTodo) error {
	db := store.db
	if filter.Search != "" {
		db = db.Where("title LIKE ? or description Like ? ", fmt.Sprintf("%%%s%%", filter.Search), fmt.Sprintf("%%%s%%", filter.Search))
	}
	if filter.SortBy != "" {
		db = db.Order(fmt.Sprintf("%s %s", filter.SortBy, filter.OrderBy))
	}

	return db.Find(&todo).Error
}

func (store *Gorm) Update(todo *model.Todo) error {
	return store.db.Save(&todo).Error
}
