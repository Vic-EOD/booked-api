package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/works", getWorks)

    router.Run("localhost:8080")
}

type work struct {
    ID string `json:"id"`
    MovieTitle string `json:"movieTitle"`
    Director string `json:"director"`
    BookTitle string `json:"bookTitle"`
    BookAuthor string `json:"id"`
    MovieReleaseYear int64 `json:"movieReleaseYear"`
}

var works = []work{
    {ID: "1", MovieTitle: "The Prestige", Director: "Chrisopher Nolan",
    BookTitle: "The Prestige", BookAuthor: "Christopher Priest", MovieReleaseYear: 2006},
    {ID: "2", MovieTitle: "Fight Club", Director: "David Fincher",
    BookTitle: "Fight Club", BookAuthor: "Chuck Palahniuk", MovieReleaseYear: 1999},
    {ID: "3", MovieTitle: "Harry Potter and the Half-Blood Prince", Director: "David Yates",
    BookTitle: "Harry Potter and the Half-Blood Prince", BookAuthor: "J.K. Rowling", MovieReleaseYear: 2009},
    {ID: "4", MovieTitle: "Blade Runner", Director: "Ridley Scott",
    BookTitle: "Do Androids Dream of Electric Sheep?", BookAuthor: "Philip K. Dick", MovieReleaseYear: 1982},
}

func getWorks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, works)
}

