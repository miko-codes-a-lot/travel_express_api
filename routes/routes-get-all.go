package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllRoutes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, routes)
}
