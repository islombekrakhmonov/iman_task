package main

import (
	"iman_task/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	api.RegisterRoutes(r)

	err := r.Run(":9080")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}


