package main

import (
	"fmt"
	_ "rgibaiev-golang-excercise/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	_ "time"
	log "github.com/Sirupsen/logrus"
	"os"
)

func main() {
	//db, err := gorm.Open("mysql", "golang_root:golang_root@/golang_exercise")
	_, err := gorm.Open("mysql", "golang_root:golang_root@/golang_exercise")
	if err != nil {
		fmt.Print("Couldn't connect to the DB")
		return
	}

	log_file, err := os.OpenFile("logrus.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(log_file)
	defer log_file.Close()


	log.WithFields(log.Fields{
		"author": "Ruslan",
		}).Info("Some log message")

	// Create table
	//db.CreateTable(&models.Task{})

	// Drop table
	//db.DropTable(&models.Task{})

	// Create new record
	//new_user := models.User{Login: "Ruslan"}
	//db.Create(&new_user)
	//new_task := models.Task{
	//	Title: "MyTitle322",
	//	Description: "MyDescription333",
	//	Priority: 1,
	//	IsDeleted: false,
	//	IsCompleted: false,
	//}
	//db.Create(&new_task)

	// Query and Update record
	//var firstUser models.User
	//db.First(&firstUser)
	//db.Model(&firstUser).Update("Login", "Olha")

	// Delete record
	//var firstUser models.User
	//db.First(&firstUser)
	//db.Delete(&firstUser)
	//fmt.Printf()


}
