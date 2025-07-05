package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func upsert(newRoute route, id *int) *route {
	if id == nil {
		routes = append(routes, newRoute)
		return &newRoute
	}

	for i, r := range routes {
		if r.ID == *id {
			newRoute.ID = *id
			routes[i] = newRoute

			return &newRoute
		}
	}

	return nil
}

func upsertRoute(c *gin.Context) {
	var newRoute route
	var id int

	if err := c.BindJSON(&newRoute); err != nil {
		return
	}

	var status = http.StatusOK
	if c.Param("id") == "" {
		newRoute = *upsert(newRoute, nil)
		status = http.StatusCreated
	} else {
		id, _ = strconv.Atoi(c.Param("id"))
		newRoute = *upsert(newRoute, &id)
	}

	c.IndentedJSON(status, newRoute)
}
