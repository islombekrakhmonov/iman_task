package api

import (
	"iman_task/api/handler"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(r *gin.Engine) {
	
	r.GET("/generate-token", handler.GenerateToken)
	r.Use(handler.JwtMiddleware)
	r.GET("/days", handler.GetDaysLeft)

}