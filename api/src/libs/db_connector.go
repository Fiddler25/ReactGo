package libs

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
)

func DBConnector() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME"))
	log.Println("Connected to mysql.")

	if err != nil {
		log.Fatal(err)
	}

	return db
}