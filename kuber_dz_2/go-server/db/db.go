package db

import (
	"database/sql"

	"github.com/yonesko/microservice-2021-03/kuber_dz_2/go-server/go"
)

type Repo interface {
	SaveUser(user swagger.User) error
	FindUser(id int64) ([]swagger.User, error)
}

type PgRepo struct {
	Db *sql.DB
}

func (p PgRepo) SaveUser(user swagger.User) error {
	panic("implement me")
}

func (p PgRepo) FindUser(id int64) ([]swagger.User, error) {
	rows, err := p.Db.Query("select name from users")
	if err != nil {
		return nil, err
	}
	return parseUsers(rows)
}

func parseUsers(rows *sql.Rows) ([]swagger.User, error) {
	return nil, nil
}
