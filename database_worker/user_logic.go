package database_worker

type TestUser struct {
	Id      string `json:"id"`
	Name    string
	Age     int
	Address string
	Car     string
}

func SetUser(id string, Name string, Age int, Address string, Car string) TestUser {
	var newUser TestUser
	newUser.Id = id
	newUser.Name = Name
	newUser.Age = Age
	newUser.Address = Address
	newUser.Car = Car
	return newUser
}

func GetUser() {

}
