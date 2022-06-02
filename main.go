package main

import (
	// "github.com/gin-gonic/gin"
	// "github.com/Vic-EOD/booked-api/controllers"
	"database/sql"
	"fmt"
	"log"

	"github.com/Vic-EOD/booked-api/db"
	"github.com/Vic-EOD/booked-api/models"
)

var db2 *sql.DB

func main() {
    // router := gin.Default()
    // router.GET("/works", controllers.GetWorks)

    db2 = db.Connect()
    works, err := worksByAuthor("Rowling, J. K.")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Works found: %v\n", works)
    //router.Run("localhost:8080")
}

func worksByAuthor(name string) ([]models.Work, error) {
    var works []models.Work

    rows, err := db2.Query("SELECT * FROM works WHERE book_author LIKE ?", name)
    if err != nil {
        return nil, fmt.Errorf("worksByAuthor %q: %v", name, err)
    }

    defer rows.Close()

    for rows.Next() {
        var work models.Work
        if err := rows.Scan(&work.ID, &work.MovieTitle, &work.BookTitle, &work.MovieReleaseYear, &work.BookAuthor); err != nil {
            return nil, fmt.Errorf("worksByAuthor %q: %v", name, err)
        }
        works = append(works, work)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("worksByAuthor %q: %v", name , err)
    }

    return works, nil
}
