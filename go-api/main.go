package main

import (
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("mysql", "user:password@tcp(mariadb:3306)/ecommerce")
    if err != nil {
        log.Fatal(err)
    }

    router := gin.Default()

    router.GET("/users", getUsers)

    router.Run(":8080")
}

func getUsers(c *gin.Context) {
    rows, err := db.Query("SELECT UserID, UserName, Email FROM Users")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []map[string]interface{}
    for rows.Next() {
        var userID int
        var userName, email string
        if err := rows.Scan(&userID, &userName, &email); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        user := map[string]interface{}{
            "UserID":   userID,
            "UserName": userName,
            "Email":    email,
        }
        users = append(users, user)
    }

    c.JSON(200, users)
}
