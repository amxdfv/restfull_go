package database_worker

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//_ "github.com/colinking/go-sqlite3-native"
	"github.com/google/uuid"
	_ "github.com/iamacarpet/go-sqlite3-dynamic"
)

func createTable(db *sql.DB) {
	users_table := `CREATE TABLE IF NOT EXISTS users (
        "id" CHAR,
        "Name" CHAR,
        "Age" INT,
        "Address" CHAR,
        "Car" CHAR);`
	query, _ := db.Prepare(users_table)
	query.Exec()
	fmt.Println("Table created successfully!")

}

func GetUserFromDB(id string) ([]byte, error) {
	db := returnDB()
	//var err error
	var YUser TestUser

	getUserRequest := "select * from users where id =?"
	res, err := db.Query(getUserRequest, id)

	defer res.Close()

	res.Next()
	err = res.Scan(&YUser.Id, &YUser.Name, &YUser.Age, &YUser.Address, &YUser.Car)
	resp_body, err := json.Marshal(YUser)
	if YUser.Id == "" {
		return nil, err
	}
	return resp_body, err
}

func AddUserToDB(req_body []byte) (string, error) {
	db := returnDB()
	var UUuser TestUser
	err := json.Unmarshal(req_body, &UUuser)
	addUserRequest := "INSERT INTO users (id, Name, Age, Adress, Car) VALUES (?, ?, ?, ?, ?)"
	id := uuid.New().String()
	_, err = db.Exec(addUserRequest, id, UUuser.Name, UUuser.Age, UUuser.Address, UUuser.Car)
	return id, err
}

func DeleteUserFromDB(req_body []byte) (string, error) {
	db := returnDB()
	var DUser TestUser
	var message string
	err := json.Unmarshal(req_body, &DUser)
	searchRes, err := GetUserFromDB(DUser.Id)
	if searchRes == nil {
		message = "Пользователь не найден"
		return message, err
	}
	delUserRequest := "delete from users where id =?"
	res, err := db.Exec(delUserRequest, DUser.Id)
	deleted_rows, err := res.RowsAffected()
	message = string(deleted_rows)
	return message, err

}

func UpdateUserFromDB(req_body []byte) (string, error) {
	db := returnDB()
	var message string
	var UpdUser TestUser
	var OrigUser TestUser
	err := json.Unmarshal(req_body, &UpdUser)
	rawUser, err := GetUserFromDB(UpdUser.Id) // поищем запись для начала
	if rawUser == nil {
		message = "Пользователь не найден"
		return message, err
	}
	err = json.Unmarshal(rawUser, &OrigUser)
	if UpdUser.Name == "" { // это если каких-то параметров нет, то возьмем старые
		UpdUser.Name = OrigUser.Name
	}
	if UpdUser.Age == 0 {
		UpdUser.Age = OrigUser.Age
	}
	if UpdUser.Address == "" {
		UpdUser.Address = OrigUser.Address
	}
	if UpdUser.Car == "" {
		UpdUser.Car = OrigUser.Car
	}

	updUserRequest := "UPDATE users SET Name=?,Age=?, Adress=?,Car=? WHERE id = ?"
	_, err = db.Exec(updUserRequest, UpdUser.Name, UpdUser.Age, UpdUser.Address, UpdUser.Car, UpdUser.Id)
	message = "Пользователь успешно обновлен"
	return message, err

}

func returnDB() *sql.DB {
	database, _ := sql.Open("sqlite3", "databsase.db")
	return database
}
