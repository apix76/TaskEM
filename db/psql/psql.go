package psql

import (
	"TaskEM/entities"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type DbAccess struct {
	db *sql.DB
}

/*
db:
	table users:
		id text
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
	sql := sqlbuilder.CreateTable("users").IfNotExists().
		Define("id", "TEXT", "NOT NULL", "PRIMARY KEY").
		Define("name", "TEXT", "NOT NULL").
		Define("surname", "TEXT", "NOT NULL").
		Define("patronymic", "TEXT").
		Define("age", "INTEGER", "NOT NULL").
		Define("race", "TEXT", "NOT NULL").
		Define("gender", "TEXT", "NOT NULL").String()

	_, err = Db.db.Exec(sql)
	return Db, err
}

func (db *DbAccess) Update(user entities.User) error {
	sqlRequest := sqlbuilder.Update("users")
	if user.Name != "" {
		sqlRequest.Set(sqlRequest.Assign("name", user.Name))
	}
	if user.Age != 0 {
		sqlRequest.Set(sqlRequest.Assign("age", user.Age))
	}
	if user.Race != "" {
		sqlRequest.Set(sqlRequest.Assign("race", user.Race))
	}
	if user.Gender != "" {
		sqlRequest.Set(sqlRequest.Assign("gender", user.Gender))
	}
	if user.Patronymic != "" {
		sqlRequest.Set(sqlRequest.Assign("patronymic", user.Patronymic))
	}
	if user.Surname != "" {
		sqlRequest.Set(sqlRequest.Assign("surname", user.Surname))
	}

	sqlRequest.Where("id = ", user.Id)

	query, args := sqlRequest.Build()

	_, err := db.db.Exec(query, args...)
	return err
}

func (db *DbAccess) Add(user entities.User) error {
	sqlRequest := sqlbuilder.InsertInto("users").
		Cols("id", "name", "surname", "patronymic", "age", "gender", "race").
		Values(user.Id, user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.Race)

	query, args := sqlRequest.Build()
	if _, err := db.db.Exec(query, args...); err != nil {
		return err
	}
	return nil
}

func (db *DbAccess) Get(cond *entities.Cond) ([]entities.User, error) {
	var users []entities.User

	sqlRequest := sqlbuilder.Select("*").From("users")

	if cond.Id != nil {
		sqlRequest.Where(sqlRequest.Equal("id", *cond.Id))
	}
	if cond.Name != nil {
		sqlRequest.Where(sqlRequest.Equal("name", *cond.Name))
	}
	if cond.Surname != nil {
		sqlRequest.Where(sqlRequest.Equal("surname", *cond.Surname))
	}
	if cond.Race != nil {
		sqlRequest.Where(sqlRequest.Equal("race", *cond.Race))
	}
	if cond.Gender != nil {
		sqlRequest.Where(sqlRequest.Equal("gender", *cond.Gender))
	}
	if cond.Patronymic != nil {
		sqlRequest.Where(sqlRequest.Equal("patronymic", *cond.Patronymic))
	}
	if cond.AgeLt != nil {
		sqlRequest.Where(sqlRequest.LE("age", cond.AgeLt))
	}
	if cond.AgeGt != nil {
		sqlRequest.Where(sqlRequest.GT("age", cond.AgeGt))
	}

	query, args := sqlRequest.Build()

	rows, err := db.db.Query(query, args...)
	defer rows.Close()

	for rows.Next() {
		var user entities.User

		if err = rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Gender, &user.Race); err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	return users, err
}

func (db *DbAccess) Delete(id string) error {
	_, err := db.db.Exec("delete from users where id = $1", id)
	return err
}
