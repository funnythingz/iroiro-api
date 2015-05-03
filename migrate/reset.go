package main

import (
	"../db"
	"../tables"
	_ "github.com/k0kubun/pp"
	"log"
)

func main() {
	db.DbLoad()
	//dropDatabase()
	//createDatabase()
	dropTable()
	createTable()
}

func dropDatabase() {
	log.Println(db.Dbmap.Exec("drop database iroiro;"))
}

func createDatabase() {
	log.Println(db.Dbmap.Exec("create database iroiro;"))
}

func dropTable() {
	log.Println(db.Dbmap.DropTableIfExists(&tables.Iro{}))
	log.Println(db.Dbmap.DropTableIfExists(&tables.Iro{}))
}

func createTable() {
	log.Println(db.Dbmap.CreateTable(&tables.Iro{}))
}
