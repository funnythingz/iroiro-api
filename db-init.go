package main

import (
	"./db"
	"fmt"
	"log"
	"os"
)

func dbLoad() {
	env := "development"
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}

	if env == "production" {
		db.DbConnect("production")
	} else {
		db.DbConnect("development")
	}

	log.Println(fmt.Sprintf("mode: %s", env))
}
