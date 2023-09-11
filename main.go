package main

import (
	"net/http"
	"sqlite-test/server_worker"
)

func returnGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write(server_worker.GetReturn(r))
	case http.MethodPost:
		w.Write(server_worker.CreateReturn(r))
	case http.MethodPut:
		w.Write(server_worker.UpdateReturn(r))
	case http.MethodDelete:
		w.Write(server_worker.DeleteReturn(r))
	}
}

func main() {
	http.HandleFunc("/oslic", returnGet)
	http.ListenAndServe(":8080", nil)
}
