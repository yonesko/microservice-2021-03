package swagger

import (
	"database/sql"
)

//CREATE TABLE users ( id serial PRIMARY KEY, username VARCHAR ( 50 )  NOT NULL, firstName VARCHAR ( 50 ) NOT NULL, lastName VARCHAR ( 50 ) NOT NULL, email VARCHAR ( 255 )  NOT NULL, phone VARCHAR ( 255 )  NOT NULL);
type Repo interface {
	SaveUser(user User) (int, error)
	FindUser(id int64) ([]User, error)
	DeleteUser(id int64) error
}

type PgRepo struct {
	Db *sql.DB
}

func (p PgRepo) SaveUser(user User) (int, error) {
	lastInsertId := 0
	rows := p.Db.QueryRow("insert into users(username, firstName, lastName, email, phone) values ($1,$2,$3,$4,$5) returning id",
		user.Username, user.FirstName, user.LastName, user.Email, user.Phone)
	err := rows.Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (p PgRepo) FindUser(id int64) ([]User, error) {
	rows, err := p.Db.Query("select * from users where id=$1", id)
	if err != nil {
		return nil, err
	}
	return parseUsers(rows)
}

func (p PgRepo) DeleteUser(id int64) error {
	panic("implement me")
}

func parseUsers(rows *sql.Rows) ([]User, error) {
	var r []User
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		r = append(r, user)
	}
	return r, nil
}
