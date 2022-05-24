package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Vic-EOD/booked-api/models"
)

func GetWorks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, models.Works)
}
