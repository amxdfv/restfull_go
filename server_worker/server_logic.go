package server_worker

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"sqlite-test/database_worker"
)

func GetReturn(request *http.Request) ([]byte, int) {
	result, err := database_worker.GetUserFromDB(request.URL.Query().Get("id"))
	code := 200
	b := []byte("Пользователь не найден")
	if err != nil {
		b = []byte(err.Error())
		return b, 418
	}
	if result == nil {
		return b, code
	} else {
		return result, code
	}
}

func CreateReturn(request *http.Request) ([]byte, int) {
	result, err := database_worker.AddUserToDB(readBody(request))
	b := []byte(result)
	if err != nil {
		b = []byte(err.Error())
		return b, 418
	}
	return b, 200
}

func DeleteReturn(request *http.Request) ([]byte, int) {
	result, err := database_worker.DeleteUserFromDB(readBody(request))
	b := []byte(result)
	if err != nil {
		b = []byte(err.Error())
		return b, 418
	}
	return b, 200
}

func UpdateReturn(request *http.Request) ([]byte, int) {
	result, err := database_worker.UpdateUserFromDB(readBody(request))
	b := []byte(result)
	if err != nil {
		b = []byte(err.Error())
		return b, 418
	}
	return b, 200
}

func readBody(request *http.Request) []byte {
	buf := make([]byte, 256)
	request.Body.Read(buf)
	buf = bytes.Trim(buf, "\x00")
	return buf
}

func SecurityCheck(request *http.Request) bool {
	user, pass, ok := request.BasicAuth()
	h := md5.Sum([]byte(pass))
	hpass := hex.EncodeToString(h[:])
	fmt.Println(hpass)
	if user == "oslic" && hpass == "418f37a7c234f270aa952991f0b90d10" && ok == true { /// TODO: потом тут сделаю шифрование
		return true
	} else {
		return false
	}
}
