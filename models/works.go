package models

type Work struct {
    ID string `json:"id"`
    MovieTitle string `json:"movieTitle"`
    BookTitle string `json:"bookTitle"`
    BookAuthor string `json:"bookAuthor"`
    MovieReleaseYear int64 `json:"movieReleaseYear"`
}

