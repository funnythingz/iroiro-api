package main

import (
	"../db"
	"../tables"
	"log"
)

func main() {
	db.DbLoad()
	migrate()
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&tables.Iro{}))
}
