package handler

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func separatePortFromHost(host string, isHTTPS bool) (string, string) {
	host, port, err := net.SplitHostPort(host)

	if err == nil {
		return port, host
	}

	if isHTTPS {
		return "443", host
	}

	return "80", host
}

func GetPortFromHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		port, hostname := separatePortFromHost(c.Request.Host, c.Request.TLS != nil)

		c.JSON(http.StatusOK, gin.H{
			"message":   hostname + ":" + port + " and is healthy",
			"status":    "success",
			"code":      http.StatusOK,
			"client_ip": c.ClientIP(),
		})
	}
}

func RouteNotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":      "route not found",
			"request_id": c.GetString("request_id"),
		})
	}
}
