package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)

type User struct {
	gorm.Model
	Name  string
	Birth string
	Email string
	Tell  string
}

var DB *gorm.DB

func DBinit() {
	var err error
	DB, err = bootDB(
		os.Getenv("APP_DB_ENDPOINT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
	)
	if err != nil {
		log.Print(err)
		return
	}

	DB.AutoMigrate(&User{})
}

func Insert(user User) {
	var err error
	DB, err = bootDB(
		os.Getenv("APP_DB_ENDPOINT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
	)
	if err != nil {
		log.Print(err)
		return
	}

	DB.Create(&user)

	defer DB.Close()
}

func bootDB(host, user, pass string) (*gorm.DB, error) {
	var err error

	driver := "mysql"
	protocol := "tcp"
	port := 3306
	name := "study_golang"
	args := "?charset=utf8&parseTime=True&loc=Local"

	con, err := gorm.Open(driver,
		fmt.Sprintf("%s:%s@%s([%s]:%d)/%s%s", user, pass, protocol, host, port, name, args),
	)
	if err != nil {
		return nil, err
	}

	con.DB().SetConnMaxLifetime(time.Second * 10)

	err = con.DB().Ping()
	if err != nil {
		return nil, err
	}

	return con, nil
}
