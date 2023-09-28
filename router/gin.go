package router

import (
	"github.com/Jakkrittkt/hugeman-assignment-golang-todo/todo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// var validate = validator.New()

type MyContext struct {
	*gin.Context
	*validator.Validate
}

func (c *MyContext) Bind(v interface{}) error {

	if err := c.Context.ShouldBindJSON(v); err != nil {
		return err
	}

	return c.Validate.Struct(v)
}

func (c *MyContext) JSON(statuscode int, v interface{}) {
	c.Context.JSON(statuscode, v)
}

func (c *MyContext) Param(key string) string {
	return c.Context.Param(key)
}

func (c *MyContext) Query(key string) string {
	return c.Context.Query(key)
}

func NewGinHandler(handler func(todo.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validate := validator.New()
		if err := validate.RegisterValidation("imageBase64", todo.CustomValidateImageBase64); err != nil {
			panic(err)
		}
		handler(&MyContext{Context: ctx, Validate: validate})
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
	}
	r.Use(cors.New(config))

	return &MyRouter{r}
}

func (r *MyRouter) POST(path string, handler func(todo.Context)) {
	r.Engine.POST(path, NewGinHandler(handler))
}

func (r *MyRouter) GET(path string, handler func(todo.Context)) {
	r.Engine.GET(path, NewGinHandler(handler))
}

func (r *MyRouter) PUT(path string, handler func(todo.Context)) {
	r.Engine.PUT(path, NewGinHandler(handler))
}
