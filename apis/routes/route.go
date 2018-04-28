package routes

import (
	"fmt"
	"gincheese/apis/handles/sessions"
	"gincheese/apis/handles/todos"
	"gincheese/apis/handles/users"
	"gincheese/apis/middleware"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
)

func RouteEngine(r *gin.Engine) *gin.Engine {

	logfile, fileErr := os.Create("/var/log/server.log")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	gin.DefaultWriter = logfile
	// r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.POST("/sessions", sessions.Create)
	r.POST("/users", users.Create)
	r.GET("/valid_email", users.SendRegistCodeToEmail)

	authorized := r.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/todos", todos.List)
		authorized.POST("/todos", todos.Create)
		authorized.GET("/todos/:id", todos.Show)
		authorized.DELETE("/todos/:id", todos.Delete)

		authorized.POST("/todos/:id/dones", todos.CreateDone)
	}

	return r
}
