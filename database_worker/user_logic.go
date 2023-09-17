package database_worker

import "strings"

type TestUser struct {
	Id      string `json:"id"`
	Name    string
	Age     int
	Address string
	Car     string
}

func RemoveDot(Usr TestUser) TestUser {
	Usr.Id = strings.ReplaceAll(Usr.Id, ";", "")
	Usr.Name = strings.ReplaceAll(Usr.Name, ";", "")
	Usr.Address = strings.ReplaceAll(Usr.Address, ";", "")
	Usr.Car = strings.ReplaceAll(Usr.Car, ";", "")
	return Usr
}
