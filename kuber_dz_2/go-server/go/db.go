package swagger

import (
	"database/sql"
)

//CREATE TABLE users ( id serial PRIMARY KEY, username VARCHAR ( 50 )  NOT NULL, firstName VARCHAR ( 50 ) NOT NULL, lastName VARCHAR ( 50 ) NOT NULL, email VARCHAR ( 255 )  NOT NULL, phone VARCHAR ( 255 )  NOT NULL);
type Repo interface {
	SaveUser(user User) error
	FindUser(id int64) ([]User, error)
}

type PgRepo struct {
	Db *sql.DB
}

func (p PgRepo) SaveUser(user User) error {
	_, err := p.Db.Query("insert into users(username, firstName, lastName, email, phone) values ($1,$2,$3,$4,$5)",
		user.Username, user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (p PgRepo) FindUser(id int64) ([]User, error) {
	rows, err := p.Db.Query("select * from users where id=$1", id)
	if err != nil {
		return nil, err
	}
	return parseUsers(rows)
}

func parseUsers(rows *sql.Rows) ([]User, error) {
	return nil, nil
}
