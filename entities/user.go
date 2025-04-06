package entities

type User struct {
	Id         string
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	Race       string
}

type UserRequest struct {
	Id         string
	Name       string
	Surname    string
	Patronymic string
}
