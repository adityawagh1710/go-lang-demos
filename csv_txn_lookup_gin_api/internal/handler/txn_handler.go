package handler

import (
	"net/http"

	"csv-txn-lookup-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func GetTxnHandler(svc *service.TxnService) gin.HandlerFunc {

	return func(c *gin.Context) {
		txnID := c.Param("id")

		txn, err := svc.Lookup(txnID)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, txn)
	}
}
