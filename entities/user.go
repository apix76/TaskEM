package entities

type User struct {
	Id         int
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	Race       string
}

type UserRequest struct {
	Name       string
	Surname    string
	Patronymic string
}
