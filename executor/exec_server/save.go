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
func insertTestData(db *gorm.DB) {
	db.Create(&Result{
		UID:         30000,
		Instruction: "test_inst1",
		Resources: []Resource{
			{
				Name:    "memory[ffff]",
				Vmiss:   58,
				Student: 0,
			},
			{
				Name:    "memory[ff1f]",
				Vmiss:   59,
				Student: 1,
			},
			{
				Name:    "memory[ff00]",
				Vmiss:   60,
				Student: 2,
			},
		}})
	db.Create(&Result{
		UID:         30010,
		Instruction: "test_inst2",
		Resources: []Resource{
			{
				Name:    "memory[0000]",
				Vmiss:   61,
				Student: 3,
			},
			{
				Name:    "memory[0001]",
				Vmiss:   62,
				Student: 4,
			},
		}})
	db.Create(&Result{
		UID:         30020,
		Instruction: "test_inst3",
		Resources: []Resource{
			{
				Name:    "memory[4441]",
				Vmiss:   63,
				Student: 5,
			},
		}})
}

func saveResult(db *gorm.DB, userid string) {
	result := jsonDecode("out/" + userid + ".json")
	fmt.Println("id : ", result.UID)
	fmt.Println("instruction : ", result.Instruction)
	for _, r := range result.Resources {
		fmt.Println("name : ", r.Name, ",vmiss : ", r.Vmiss, ",student : ", r.Student)
	}
	id, _ := strconv.Atoi(userid)
	db.Create(&Result{
		UID:         id,
		Instruction: result.Instruction,
		Resources:   result.Resources,
	})
}

func SaveMysql(userid string) {
	db := gormConnect()
	defer db.Close()
	dbInit(db)
	saveResult(db, userid)
}
