package main

import "log"

func init() {
	db := gormConnect()
	dbInit(db)
}

func main() {
	const user = "00000000" //test
	Exec(user)
	if err := SaveMysql(user); err != nil {
		log.Fatal(err)
	}
}
