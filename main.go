package main

import (
	"github.com/AlmasOrazgaliev/assignment3/controller"
	"github.com/AlmasOrazgaliev/assignment3/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(
		postgres.Open("host=dpg-ci04ksd269v5qbksrnmg-a port=5432 user=kbtugophers_user password=3Hql6UTpXqgK26RMsYfUTeV4fBBv8R4M dbname=kbtugophers_db sslmode=disable"),
		&gorm.Config{})
	db.AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatal(err)
	}
	contr := controller.NewController(db)
	controller.Start(contr)
}
