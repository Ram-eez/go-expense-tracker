package main

import "expensetracker/internal"

func main() {
	// router := gin.Default()

	// api.RegisterRoutes(router)

	// router.Run(":8080")

	internal.ParseMessage("debited by 900.00 on date 9feb24 trd to rameez")
}
