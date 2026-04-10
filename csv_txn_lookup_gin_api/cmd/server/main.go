package main

import (
	"csv-txn-lookup-gin-api/internal/handler"
	"csv-txn-lookup-gin-api/internal/middleware"
	"csv-txn-lookup-gin-api/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.RequestIDMiddleware())
	server.Use(middleware.ErrorHandlerMiddleware())
	server.Use(middleware.SimpleLogger())
	server.Use(middleware.AuthMiddleware())
	server.NoRoute(handler.RouteNotFoundHandler())
	server.GET("/", handler.GetPortFromHost())

	v1 := server.Group("/api/v1")

	router.TxnLookupRoutes(v1)

	server.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	server.RunTLS(":443", "cert.pem", "key.pem")
	server.Run(":8080")
}
