package api

import (
	"expensetracker/internal"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTransaction(c *gin.Context) {

	msg, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newMessage := internal.ParseMessage(string(msg))

	if err := newMessage.CreateTransaction(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "posted"})
}
