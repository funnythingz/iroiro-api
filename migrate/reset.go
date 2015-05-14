package main

import (
	"../db"
	"../tables"
	_ "github.com/k0kubun/pp"
	"log"
)

func main() {
	db.DbLoad()
	dropTable()
	createTable()
}

func dropTable() {
	log.Println(db.Dbmap.DropTableIfExists(&domain.Iro{}))
	log.Println(db.Dbmap.DropTableIfExists(&domain.Color{}))
}

func createTable() {
	log.Println(db.Dbmap.CreateTable(&domain.Iro{}))
	log.Println(db.Dbmap.CreateTable(&domain.Color{}))
}
