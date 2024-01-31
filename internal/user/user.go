package user

type User struct {
	ID             int
	Name           string
	Email          string
	BirthDate      string
	Phone          string
	DocumentNumber string
	Address        Address
}

type Address struct {
	Street     string
	Number     string
	Complement string
	City       string
	Country    string
	State      string
	ZipCode    string
}
