package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Vic-EOD/booked-api/models"
    "github.com/Vic-EOD/booked-api/db"
    "database/sql"
    "log"
)

var db2 *sql.DB

func GetWorks(c *gin.Context) {
    var works []models.Work
    db2 = db.Connect();

    rows, err := db2.Query("SELECT * FROM works ORDER BY RAND() LIMIT 10")
    if err != nil {
        log.Fatalf("GetWorks: %v", err)
    }

    defer rows.Close()

    for rows.Next() {
        var work models.Work
        if err := rows.Scan(&work.ID, &work.MovieTitle, &work.BookTitle, &work.MovieReleaseYear, &work.BookAuthor); err != nil {
            log.Fatalf("GetWorks: %v", err)
        }
        works = append(works, work)
    }
    if err := rows.Err(); err != nil {
        log.Fatalf("GetWorks: %v", err)
    }

    c.IndentedJSON(http.StatusOK, works)
}
