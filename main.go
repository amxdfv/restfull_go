package main

import (
	"net/http"
	"sqlite-test/server_worker"
)

func returnGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		payload, code := server_worker.GetReturn(r)
		w.WriteHeader(code)
		w.Write(payload)
	case http.MethodPost:
		payload, code := server_worker.CreateReturn(r)
		w.WriteHeader(code)
		w.Write(payload)
	case http.MethodPut:
		payload, code := server_worker.UpdateReturn(r)
		w.WriteHeader(code)
		w.Write(payload)
	case http.MethodDelete:
		payload, code := server_worker.DeleteReturn(r)
		w.WriteHeader(code)
		w.Write(payload)
	}
}

func main() {
	http.HandleFunc("/oslic", returnGet)
	http.ListenAndServe(":8080", nil)
}
