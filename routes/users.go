package routes

import "github.com/gin-gonic/gin"
import "study_go/service"

func Users(r *gin.Engine) {
	routes := r.Group("/users")
	routes.GET("/list", service.List)
}
