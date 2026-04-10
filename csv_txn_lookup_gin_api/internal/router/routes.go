package router

import (
	"csv-txn-lookup-gin-api/internal/handler"
	"csv-txn-lookup-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func TxnLookupRoutes(rg *gin.RouterGroup) {

	svc := service.NewTxnService()

	txn := rg.Group("/txn")
	{
		txn.GET("/:id", handler.GetTxnHandler(svc))
	}
}
