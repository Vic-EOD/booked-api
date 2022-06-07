package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Vic-EOD/booked-api/controllers"
    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())
    router.GET("/works", controllers.GetWorks)

    router.Run("localhost:8080")
}

