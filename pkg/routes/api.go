package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/rest-api-go/pkg/handler/user"
)

// Routes struct
type Routes struct {
}

// StartGin will start the router
func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", welcome)
		api.GET("/users", user.GetAllUsers)
		api.POST("/users", user.CreateUser)
	}
	r.Run(":8000")
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to my API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}
