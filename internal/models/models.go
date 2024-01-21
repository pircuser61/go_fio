package models

type Person struct {
	Id          uint32
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

type Filter struct {
	Gender string
	Age    int32
	Order  string
	Offset uint64
	Limit  uint64
}
