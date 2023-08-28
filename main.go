package main

import (
	"fmt"
	"net/http"
)

func returnGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Пока не готово, но скоро будет"))
	case http.MethodPost:
		buf := make([]byte, 256)
		r.Body.Read(buf)
		fmt.Println(string(buf))
		bodyResp := string(buf)
		w.Write([]byte(bodyResp))
	case http.MethodPut:
		w.Write([]byte("Пока не готово, но скоро будет"))
	case http.MethodDelete:
		w.Write([]byte("Пока не готово, но скоро будет"))
	}
}

func main() {
	//database_worker.WorkWithDb()

	http.HandleFunc("/oslic", returnGet)
	http.ListenAndServe(":8080", nil)
}
