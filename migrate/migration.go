package main

import (
	"../db"
	"../domain"
	"log"
)

func main() {
	db.Connect()
	migrate()
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&domain.Iro{}))
	log.Println(db.Dbmap.AutoMigrate(&domain.Color{}))
}
