package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Vic-EOD/booked-api/controllers"
	//"database/sql"
	//"fmt"
	//"log"

	//"github.com/Vic-EOD/booked-api/db"
	//"github.com/Vic-EOD/booked-api/models"
)

func main() {
    router := gin.Default()
    router.GET("/works", controllers.GetWorks)

    router.Run("localhost:8080")
}

