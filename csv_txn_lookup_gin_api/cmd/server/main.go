package main

import (
	"log"

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

	if err := server.SetTrustedProxies([]string{"127.0.0.1", "::1"}); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	// Try TLS first, fall back to plain HTTP
	if err := server.RunTLS(":443", "cert.pem", "key.pem"); err != nil {
		log.Printf("TLS unavailable (%v), falling back to HTTP on :8080", err)
		if err := server.Run(":8080"); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}
}
