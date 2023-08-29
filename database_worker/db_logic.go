package database_worker

import (
	"database/sql"
	"encoding/json"
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
        "Address" CHAR,
        "Car" CHAR);`
	query, err := db.Prepare(users_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	fmt.Println("Table created successfully!")

}

func GetUserFromDB(id string) []byte {
	db := returnDB()
	var YUser TestUser

	getUserRequest := "select * from users where id =?"
	res, err := db.Query(getUserRequest, id)
	if err != nil {
		log.Fatal(err)
	}
	res.Next()
	res.Scan(&YUser.Id, &YUser.Name, &YUser.Age, &YUser.Address, &YUser.Car)
	resp_body, err := json.Marshal(YUser)
	return resp_body
}

func AddUserToDB(req_body []byte) string {
	db := returnDB()
	var UUuser TestUser
	err := json.Unmarshal(req_body, &UUuser)
	fmt.Println(err)

	addUserRequest := "INSERT INTO users (id, Name, Age, Adress, Car) VALUES (?, ?, ?, ?, ?)"
	id := uuid.New().String()
	db.Exec(addUserRequest, id, UUuser.Name, UUuser.Age, UUuser.Address, UUuser.Car)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return id
}

func returnDB() *sql.DB {
	database, _ := sql.Open("sqlite3", "databsase.db")
	return database
}
