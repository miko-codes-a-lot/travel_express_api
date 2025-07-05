package routes

import "github.com/gin-gonic/gin"

func RouteRoutes(r *gin.Engine) {
	r.GET("/routes", getAllRoutes)
	r.GET("/routes/:id", getOneRoute)
	r.POST("/routes", upsertRoute)
	r.PUT("/routes/:id", upsertRoute)
}
