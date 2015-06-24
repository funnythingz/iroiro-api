package main

import (
	"../db"
	"../mapper"
	"fmt"
	"log"
	"os"
)

func main() {
	db.Connect()

	action := "migrate"
	if len(os.Args) >= 2 {
		action = os.Args[2]
	}

	log.Println(fmt.Sprintf("mode: %s", action))

	switch {
	case action == "migrate":
		Migrate()
		return

	case action == "reset":
		Reset()
		return
	}
}

func Reset() {
	db.Dbmap.DropTableIfExists(&mapper.Iro{})
	db.Dbmap.DropTableIfExists(&mapper.Color{})
	Create()
}

func Create() {
	db.Dbmap.CreateTable(&mapper.Iro{})
	db.Dbmap.CreateTable(&mapper.Color{})
}

func Migrate() {
	db.Dbmap.AutoMigrate(&mapper.Iro{})
	db.Dbmap.AutoMigrate(&mapper.Color{})
}
