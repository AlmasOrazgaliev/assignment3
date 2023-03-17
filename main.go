package main

import (
	"github.com/AlmasOrazgaliev/assignment3/controller"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres password=alma45884 dbname=library sslmode=disable"),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	contr := controller.NewController(db)
	controller.Start(contr)
}
