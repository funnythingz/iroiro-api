package main

import (
	"../db"
	"../mapper"
	"log"
)

func main() {
	db.Connect()
	migrate()
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&mapper.Iro{}))
	log.Println(db.Dbmap.AutoMigrate(&mapper.Color{}))
}
