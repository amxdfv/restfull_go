package database_worker

import (
	"database/sql"
	"fmt"
	//_ "github.com/colinking/go-sqlite3-native"
	"github.com/google/uuid"
	_ "github.com/iamacarpet/go-sqlite3-dynamic"
	//_ "github.com/mattn/go-sqlite3"
	"log"
)

func createTable(db *sql.DB) {
	users_table := `CREATE TABLE IF NOT EXISTS users (
        "id" CHAR,
        "Name" CHAR,
        "Age" INT,
        "Adress" CHAR,
        "Car" CHAR);`
	query, err := db.Prepare(users_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	fmt.Println("Table created successfully!")

}

func GetUserFromDB(db *sql.DB, id string) {
	var idd string
	var Name string
	var Age string
	var Adress string
	var Car string

	getUserRequest := "select * from users where id =?"
	res, err := db.Query(getUserRequest, id)
	if err != nil {
		log.Fatal(err)
	}
	res.Next()
	res.Scan(&idd, &Name, &Age, &Adress, &Car)
	fmt.Println(idd + " " + Name + " " + Age + " " + Adress + " " + Car)
}

func AddUserToDB(db *sql.DB) {
	addUserRequest := "INSERT INTO users (id, Name, Age, Adress, Car) VALUES (?, ?, ?, ?, ?)"
	id := uuid.New()
	res, err := db.Exec(addUserRequest, id, "Олег", 32, "Томск", "Жигули")
	//db.Query(addUserRequest, id)
	//query, err := db.Prepare(addUserRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(id)
}

func WorkWithDb() {
	database, _ := sql.Open("sqlite3", "databsase.db")
	//	AddUserToDB(database)
	GetUserFromDB(database, "40a1709f-4006-43c4-8173-f2e2587d9274")
	//createTable(database)
}
