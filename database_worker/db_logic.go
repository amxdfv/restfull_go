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

func GetUserFromDB(id string) []byte {
	db := returnDB()
	var YUser TestUser

	getUserRequest := "select * from users where id =?"
	res, _ := db.Query(getUserRequest, id)
	defer res.Close()

	res.Next()
	res.Scan(&YUser.Id, &YUser.Name, &YUser.Age, &YUser.Address, &YUser.Car)
	resp_body, _ := json.Marshal(YUser)
	if YUser.Id == "" {
		return nil
	}
	return resp_body
}

func AddUserToDB(req_body []byte) string {
	db := returnDB()
	var UUuser TestUser
	json.Unmarshal(req_body, &UUuser)
	addUserRequest := "INSERT INTO users (id, Name, Age, Adress, Car) VALUES (?, ?, ?, ?, ?)"
	id := uuid.New().String()
	db.Exec(addUserRequest, id, UUuser.Name, UUuser.Age, UUuser.Address, UUuser.Car)
	return id
}

func DeleteUserFromDB(req_body []byte) string {
	db := returnDB()
	var DUser TestUser
	var message string
	json.Unmarshal(req_body, &DUser)
	if GetUserFromDB(DUser.Id) == nil {
		message = "Пользователь не найден"
		return message
	}
	delUserRequest := "delete from users where id =?"
	res, _ := db.Exec(delUserRequest, DUser.Id)
	deleted_rows, _ := res.RowsAffected()
	message = string(deleted_rows)
	return message

}

func UpdateUserFromDB(req_body []byte) string {
	db := returnDB()
	var message string
	var UpdUser TestUser
	var OrigUser TestUser
	json.Unmarshal(req_body, &UpdUser)
	rawUser := GetUserFromDB(UpdUser.Id) // поищем запись для начала
	if rawUser == nil {
		message = "Пользователь не найден"
		return message
	}
	json.Unmarshal(rawUser, &OrigUser)
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
	res, err := db.Exec(updUserRequest, UpdUser.Name, UpdUser.Age, UpdUser.Address, UpdUser.Car, UpdUser.Id)
	fmt.Print(res)
	fmt.Print(err)
	message = "Пользователь успешно обновлен"
	return message

}

func returnDB() *sql.DB {
	database, _ := sql.Open("sqlite3", "databsase.db")
	return database
}
