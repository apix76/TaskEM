package usecase

import (
	"TaskEM/db/psql"
	"TaskEM/entities"
	"TaskEM/framework"
	"fmt"
	"math/rand"
)

type Usecase struct {
	db psql.DbAccess
}

func NewUseCase(db *psql.DbAccess) *Usecase {
	return &Usecase{db: *db}
}

func (u *Usecase) Get(filter map[string]string) ([]entities.User, error) {
	return u.db.Get(filter)
}

func (u *Usecase) Delete(id string) error {
	return u.db.Delete(id)
}

func (u *Usecase) Post(userSet entities.UserRequest) (entities.User, error) {
	var err error

	user := entities.User{Name: userSet.Name, Surname: userSet.Surname, Patronymic: userSet.Patronymic}
	idInt := rand.Int63()
	user.Id = fmt.Sprint(idInt)
	if user.Age, err = framework.GetAge(user.Name); err != nil {
		return entities.User{}, err
	}
	if user.Gender, err = framework.GetGender(user.Name); err != nil {
		return entities.User{}, err
	}
	if user.Race, err = framework.GetRace(user.Name); err != nil {
		return entities.User{}, err
	}

	return user, u.db.Add(user)
}

func (u *Usecase) Patch(userSet entities.UserRequest) (entities.User, error) {
	var err error

	user := entities.User{Name: userSet.Name, Surname: userSet.Surname, Patronymic: userSet.Patronymic}

	if user.Age, err = framework.GetAge(user.Name); err != nil {
		return entities.User{}, err
	}
	if user.Gender, err = framework.GetGender(user.Name); err != nil {
		return entities.User{}, err
	}
	if user.Race, err = framework.GetRace(user.Name); err != nil {
		return entities.User{}, err
	}

	return u.db.Update(user)
}
