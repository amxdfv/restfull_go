package main

import (
	"bytes"
	"net/http"
	"sqlite-test/database_worker"
)

func returnGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write(database_worker.GetUserFromDB(r.URL.Query().Get("id")))
	case http.MethodPost:
		buf := make([]byte, 256)
		r.Body.Read(buf)
		buf = bytes.Trim(buf, "\x00")
		w.Write([]byte(database_worker.AddUserToDB(buf)))
	case http.MethodPut:
		buf := make([]byte, 256)
		r.Body.Read(buf)
		buf = bytes.Trim(buf, "\x00")
		w.Write([]byte(database_worker.UpdateUserFromDB(buf)))
	case http.MethodDelete:
		buf := make([]byte, 256)
		r.Body.Read(buf)
		buf = bytes.Trim(buf, "\x00")
		w.Write([]byte("Удалено пользователей: " + database_worker.DeleteUserFromDB(buf)))
	}
}

func main() {
	http.HandleFunc("/oslic", returnGet)
	http.ListenAndServe(":8080", nil)
}
