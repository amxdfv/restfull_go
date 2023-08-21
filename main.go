package main

import (
	"fmt"
	"net/http"
)

func returnGet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "oslic")
}

func main() {
	http.HandleFunc("/oslic", returnGet)
	http.ListenAndServe(":8080", nil)
}
