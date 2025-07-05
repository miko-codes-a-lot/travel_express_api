package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getOne(id int) *route {
	for _, r := range routes {
		if r.ID == id {
			return &r
		}
	}
	return nil
}

func getOneRoute(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	r := getOne(id)
	if r != nil {
		c.IndentedJSON(http.StatusOK, r)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "route not found"})
}
