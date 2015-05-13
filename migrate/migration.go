package main

import (
	"../db"
	"../tables"
	"log"
)

func main() {
	db.Connect()
	migrate()
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&tables.Iro{}))
}
