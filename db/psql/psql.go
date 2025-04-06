package psql

import (
	"TaskEM/entities"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbAccess struct {
	db *sql.DB
}

type Cond struct {
	Id           *string
	Name         *string
	AgeGt, AgeLt *int
	Surname      *string
	Race         *string
	Gender       *string
	Patronymic   *string
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

func (db *DbAccess) Update(user entities.User) error {
	sqlRequest := sqlbuilder.Update("users")
	if user.Name != "" {
		sqlRequest.Set(sqlRequest.Assign("name", user.Name),
			sqlRequest.Assign("age", user.Age),
			sqlRequest.Assign("race", user.Race),
			sqlRequest.Assign("gender", user.Gender),
		)
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

func (db *DbAccess) Get(cond *Cond) ([]entities.User, error) {
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
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

	}
	return users, err
}

func (db *DbAccess) Delete(id string) error {
	_, err := db.db.Exec("delete from users where id = $1", id)
	return err
}
