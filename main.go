package main

import (
	"fmt"
	"rgibaiev-golang-excercise/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	_ "time"
	log "github.com/Sirupsen/logrus"
	"os"
	"github.com/labstack/echo"
	"net/http"
	"github.com/labstack/echo/engine/standard"
	"strconv"
	_ "encoding/json"
)

var db *gorm.DB

func hello(c echo.Context) error {
	log.WithFields(
		log.Fields{"function": "hello",
		}).Info("Success")
	return c.String(http.StatusOK, "Hello, World!")
}

func add(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("a"))
	if err != nil {
		log.WithFields(log.Fields{
			"function": "add",
			"parameter": c.Param("a"),
			}).Fatal("Fail")
		return c.String(http.StatusBadRequest, "No parameter a passed")
		}
	log.WithFields(
		log.Fields{
			"function": "add",
			"parameter": c.Param("a"),
		}).Info("Success")
	return c.String(http.StatusOK, strconv.Itoa(i + 10))
}

func new_user(c echo.Context) error {
	u := models.User{}
	err := c.Bind(u)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid parameters passed: " + err.Error())
	} else {
		fmt.Printf(u.Login)
		return c.String(http.StatusOK, u.Login)
	}
}


func add_to_db(c echo.Context) error {
	first_name := "First Name"
	row := models.User{Login : "myLogin", FirstName : &first_name}
	db.NewRecord(row)
	db.Create(&row)
	return c.String(http.StatusOK, "OK!")
}

func main() {
	// DB setup
	var err error
	db, err = gorm.Open("mysql", "golang_root:golang_root@/golang_exercise")
	//_, err := gorm.Open("mysql", "golang_root:golang_root@/golang_exercise")
	if err != nil {
		fmt.Print("Couldn't connect to the DB")
		return
	}

	// Logging setup
	log_file, err := os.OpenFile("logrus.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(log_file)
	defer log_file.Close()

	db.CreateTable(&models.Task{})
	db.CreateTable(&models.User{})

	// Web-server setup
	e := echo.New()
	e.GET("/", hello)
	e.GET("/add/:a/", add)
	e.GET("/save/", add_to_db)
	e.POST("/users/add", new_user)
	e.Run(standard.New(":1323"))

	// Create table
	//db.CreateTable(&models.Task{})
	//db.CreateTable(&models.User{})

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
