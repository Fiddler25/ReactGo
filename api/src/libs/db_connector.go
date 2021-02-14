package libs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBConnector() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	log.Println("Connected to mysql.")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
