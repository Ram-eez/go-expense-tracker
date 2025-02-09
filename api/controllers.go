package api

import (
	"expensetracker/internal"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTransaction(c *gin.Context) {
	var newMessage internal.Message
	if err := newMessage.CreateTransaction(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "posted"})
}
