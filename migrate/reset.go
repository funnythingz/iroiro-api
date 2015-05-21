package main

import (
	"../db"
	"../mapper"
	"log"
)

func main() {
	db.Connect()
	dropTable()
	createTable()
}

func dropTable() {
	log.Println(db.Dbmap.DropTableIfExists(&mapper.Iro{}))
	log.Println(db.Dbmap.DropTableIfExists(&mapper.Color{}))
}

func createTable() {
	log.Println(db.Dbmap.CreateTable(&mapper.Iro{}))
	log.Println(db.Dbmap.CreateTable(&mapper.Color{}))
}
