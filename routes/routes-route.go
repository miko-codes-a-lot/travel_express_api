package routes

import "github.com/gin-gonic/gin"

func RouteRoutes(r *gin.Engine) {
	r.GET("/api/routes", getAllRoutes)
	r.GET("/api/routes/:id", getOneRoute)
	r.POST("/api/routes", upsertRoute)
	r.PUT("/api/routes/:id", upsertRoute)
}
