package main

import (
	"fmt"

	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/db"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/router"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo"
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo/model"
)

func main() {
	// dbGorm, err := gorm.Open(sqlite.Open("sqlite3"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// dbGorm.AutoMigrate(&model.Todo{})
	// gorm := db.NewGorm(dbGorm)
	dbJson := db.NewDbJsonFile(fmt.Sprintf("%s.json", model.Todo{}.TableName()))
	todoHandler := todo.NewTodoHandler(dbJson)

	r := router.NewMyRouter()
	r.GET("/todos", todoHandler.FindAll)
	r.GET("/todos/:id", todoHandler.FindById)
	r.POST("/todos", todoHandler.Create)
	r.PUT("/todos/:id", todoHandler.UpdateById)
	r.Run()
}
