package main

import (
	"project/internal/db"
	"project/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	// Test url ab -n 1000 -c 50 http://127.0.0.1:8080/users/1
	r.GET("/users/:id", handler.GetUser)

	r.Run(":8080")
}
