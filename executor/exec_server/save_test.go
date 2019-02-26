package main

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB

func TestMain(m *testing.M) {
	dbConn = gormConnect()
	defer dbConn.Close()
	dbInit(dbConn)
	insertTestData(dbConn)

	result := m.Run()
	os.Exit(result)
}

func TestSaveMysql(t *testing.T) {
	testUser := "00000000"
	Exec(testUser)
	if err := SaveMysql(testUser); err != nil {
		t.Errorf("Failed save mysql user data. got error=%v", err)
	}
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
