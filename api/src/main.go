package main

import (
	_ "github.com/go-sql-driver/mysql"
	"queries"
	"server"
)

func main() {
	queries.Query()

	server.Init()
}