package psql

import (
	"TaskEM/entities"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbAccess struct {
	db *sql.DB
}

/*
db:
	table users:
		id int
		name text
		surname text
		patronymic text (can be empty)
		age int
		gender text
		race text
*/

func NewDb(dsn string) (DbAccess, error) {
	var Db DbAccess
	var err error

	if Db.db, err = sql.Open("pgx", dsn); err != nil {
		return Db, err
	}
	return Db, err
}

func (db *DbAccess) Patch(user entities.User){
	sqlbuilder.Update("users")
}

func (db *DbAccess) Add(user entities.User) error {
	sqlRequest := sqlbuilder.InsertInto("users").
		Cols("id", "name", "surname", "patronymic", "age", "gender", "race").
		Values(user.Id, user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.Race)

	if _, err := db.db.Exec(sqlRequest.String()); err != nil {
		return err
	}
	return nil
}

func (db *DbAccess) Get(condition map[string]string) (string, error) {
	sqlRequest := sqlbuilder.Select("*").From("users")

	if len(condition) != 0 {
		for _, v := range condition {
			sqlRequest = sqlRequest.Where(v)
		}
	}

	row := db.db.QueryRow(sqlRequest.String())

	var URL string

	if err := row.Scan(&URL); err != nil {
		if err == sql.ErrNoRows {
			return "",
		} else {
			return "", err
		}
	}

	return URL, nil
}

func (db *DbAccess) Delete(id int) error {
	_, err := db.db.Exec("delete from users where id = $1", id)
	return err
}
