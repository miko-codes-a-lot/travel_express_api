package main

import (
	"cloudants/travel-express/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RouteRoutes(router)

	router.Run("localhost:8080")
}
