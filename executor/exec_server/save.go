package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Result struct {
	gorm.Model
	UID         int        `json:"id"`
	Instruction string     `json:"instruction"`
	Resources   []Resource `json:"resource" gorm:"foreignkey:UserRefer"`
}
type Resource struct {
	gorm.Model
	Name      string `json:"name"`
	Vmiss     int    `json:"vmiss"`
	Student   int    `json:"student"`
	UserRefer uint
}

func jsonDecode(filename string) Result {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "vmiss"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInit(db *gorm.DB) {
	db.DropTableIfExists(&Result{}, &Result{})
	db.AutoMigrate(&Result{})
	db.AutoMigrate(&Resource{})
	//insertTestData(db)
}

func saveResult(db *gorm.DB, userid string) error {
	result := jsonDecode("out/" + userid + ".json")
	fmt.Println("id : ", result.UID)
	fmt.Println("instruction : ", result.Instruction)
	for _, r := range result.Resources {
		fmt.Println("name : ", r.Name, ",vmiss : ", r.Vmiss, ",student : ", r.Student)
	}
	id, err := strconv.Atoi(userid)
	if err != nil {
		return err
	}
	if err := db.Create(&Result{
		UID:         id,
		Instruction: result.Instruction,
		Resources:   result.Resources,
	}).Error; err != nil {
		return err
	}
	return nil
}

func SaveMysql(userid string) error {
	db := gormConnect()
	defer db.Close()
	return saveResult(db, userid)
}
