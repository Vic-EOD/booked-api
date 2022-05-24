package main

import (
    "github.com/gin-gonic/gin"
    "github.com/Vic-EOD/booked-api/controllers"
    "github.com/Vic-EOD/booked-api/db"
)

func main() {
    router := gin.Default()
    router.GET("/works", controllers.GetWorks)

    db.Connect()

    router.Run("localhost:8080")
}

