package handler

import (
	"net/http"
	"project/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.ParseInt(idParam, 10, 64)

	res, err := service.GetUserFullData(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, res)
}
