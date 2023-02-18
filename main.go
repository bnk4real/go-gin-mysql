package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type ssdusers struct {
	UserId      string
	FulllNameTh string
}

func main() {

	// Connect to MySQL Database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/databaseName")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if connection is still alive
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// Call API to get users data
	r.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT userID, fullNameTH FROM tablename")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		dto := ssdusers{}

		users := []gin.H{}

		for rows.Next() {
			err := rows.Scan(&dto.UserId, &dto.FulllNameTh)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, gin.H{
				"User ID":   dto.UserId,
				"Full Name": dto.FulllNameTh})
		}
		// Return results as JSON body
		c.JSON(200, gin.H{
			"Results": users,
		})
	})

	// Run on port 8090
	r.Run(":8090")
}
