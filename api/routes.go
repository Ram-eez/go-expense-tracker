package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.POST("/transaction", PostTransaction)
	router.GET("/transaction/:trans_id", GetTransactionByID)
	router.GET("/transaction/all", GetAllTransactions)
	router.GET("/transaction/all/amount/asc", SortByAmountAsc)
	router.GET("/transaction/all/amount/desc", SortByAmountDesc)
	router.GET("/transaction/all/date/asc", SortByDateAsc)
	router.GET("/transaction/all/date/desc", SortByDateDesc)
	router.GET("/transaction/all/total", TotalExpenses)
	router.DELETE("/transaction/:trans_id", DeleteTransaction)
	router.PATCH("/transaction/:trans_id", UpdateTransaction)
}
