package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.Engine) {
	Users(r)
}
