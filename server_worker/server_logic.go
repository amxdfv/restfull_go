package server_worker

import (
	"bytes"
	"net/http"
	"sqlite-test/database_worker"
)

func GetReturn(request *http.Request) []byte {
	result, err := database_worker.GetUserFromDB(request.URL.Query().Get("id"))
	b := []byte("Пользователь не найден")
	if err != nil {
		b = []byte(err.Error())
	}
	if result == nil {
		return b
	} else {
		return result
	}
}

func CreateReturn(request *http.Request) []byte {
	result, err := database_worker.AddUserToDB(readBody(request))
	b := []byte(result)
	if err != nil {
		b = []byte(err.Error())
		return b
	}
	return b
}

func DeleteReturn(request *http.Request) []byte {
	result, err := database_worker.DeleteUserFromDB(readBody(request))
	b := []byte(result)
	if err != nil {
		b = []byte(err.Error())
		return b
	}
	return b
}

func UpdateReturn(request *http.Request) []byte {
	result, err := database_worker.UpdateUserFromDB(readBody(request))
	b := []byte(result)
	if err != nil {
		b = []byte(err.Error())
		return b
	}
	return b
}

func readBody(request *http.Request) []byte {
	buf := make([]byte, 256)
	request.Body.Read(buf)
	buf = bytes.Trim(buf, "\x00")
	return buf
}
