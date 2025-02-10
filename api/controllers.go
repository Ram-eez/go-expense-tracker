package api

import (
	"expensetracker/internal"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostTransaction(c *gin.Context) {

	msg, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMessage := internal.ParseMessage(string(msg))

	if err := newMessage.CreateTransaction(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "posted"})
}

func GetTransactionByID(c *gin.Context) {
	ID := c.Param("trans_id")
	TransID, err := strconv.ParseInt(ID, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse ID"})
		return
	}

	message := &internal.Message{Id: TransID}
	message, err = message.GetTransaction()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get the message from db"})
		return
	}

	c.JSON(http.StatusOK, message)
}
