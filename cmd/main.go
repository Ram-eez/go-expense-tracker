package main

import (
	"expensetracker/internal"
	"fmt"
)

func main() {
	// router := gin.Default()

	// api.RegisterRoutes(router)

	// router.Run(":8080")

	message := internal.ParseMessage("debited by 900.00 on date 9feb24 trd to rameez")
	fmt.Println(message)
}
