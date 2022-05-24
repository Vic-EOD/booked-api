package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
    var db *sql.DB
    cfg := mysql.Config{
        User: os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net: "tcp",
        Addr: "127.0.0.1:3306",
        DBName: "movie_books",
    }

    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

    fmt.Println("Connected!")
    return db
}