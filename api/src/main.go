package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-contrib/sessions/cookie"
	"log"
	"libs"
	. "fmt"
	"routes"
)

type Person struct {
	ID   int
	Name string
}

func main() {
	r := gin.Default()

	db := libs.DBConnector()
	defer db.Close()

	sql, err := db.Prepare("INSERT INTO users(name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	sql.Exec("testUser")

	sql, err = db.Prepare("UPDATE users SET name = ? WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}
	sql.Exec("updateUser")

	sql, err = db.Prepare("DELETE FROM users WHERE id = 2")
	if err != nil {
		log.Fatal(err)
	}
	sql.Exec()

	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.Name)

		if err != nil {
			log.Fatal(err)
		}
		Println(person.ID, person.Name)
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4000"}
	r.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
 
	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{ "hello": session.Get("hello") })
	})
	
	r.GET("/", routes.Index)
	r.POST("/login", routes.Login)

	r.Run()
}