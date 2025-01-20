
package main

import (
	"github.com/gin-gonic/gin"

	"myproject/dummy/router"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	dummyRouter.RouterInit(r)

	// Start the server on port 8080
	r.Run(":8080")
}
