package queries

import (
	"log"
	"libs"
	. "fmt"
	"viewmodels"
)

func Query() {
	db := libs.DBConnector()
	defer db.Close()
	
	query, err := db.Query("SELECT * FROM users")
	defer query.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	for query.Next() {
		var user viewmodels.User
		err := query.Scan(&user.ID, &user.Name)

		if err != nil {
			log.Fatal(err)
		}
		Println(user.ID, user.Name)
	}
}