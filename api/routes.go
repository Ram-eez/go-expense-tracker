package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.POST("/transaction", PostTransaction)
	router.GET("/transaction/:trans_id", GetTransactionByID)
	router.GET("/transaction/all", GetAllTransactions)
	router.GET("/transaction/all/asc", SortByDateAsc)
	router.DELETE("/transaction/:trans_id", DeleteTransaction)
	router.PATCH("/transaction/:trans_id", UpdateTransaction)
}
