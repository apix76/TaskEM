package entities

type Cond struct {
	Id                *string
	Name              *string
	AgeGt, AgeLt, Age *int
	Surname           *string
	Race              *string
	Gender            *string
	Patronymic        *string
}
