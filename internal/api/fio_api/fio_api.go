package fio_api

type Api interface {
	GetGender(name string) (string, error)
	GetAge(name string) (int, error)
	GetNationality(name string) (string, error)
}
