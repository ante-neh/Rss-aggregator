package database

import (
	"database/sql"
	"github.com/ante-neh/Rss-aggregator/types"
	"github.com/ante-neh/Rss-aggregator/util"
	"github.com/google/uuid"
	"time"
)

type DatabaseOperation interface {
	Createuser(id uuid.UUID, created_at time.Time, updated_at time.Time, name string) (types.User, error)
	GetUser(api_key string) (types.User, error)
}
type Postgres struct {
	DB *sql.DB
}

func (p *Postgres) Createuser(id uuid.UUID, created_at time.Time, updated_at time.Time, name string) (types.User, error) {
	stmt := "INSERT INTO users(id, created_at, updated_at, name, api_key) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at, name, api_key"

	var user types.User
	err := p.DB.QueryRow(stmt, id, created_at, updated_at, name, util.GenerateApikey()).Scan(&user.ID, &user.Created_at, &user.Updated_at, &user.Name, &user.Api_key)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (p *Postgres) GetUser(api_key string) (types.User, error) {
	stmt := "SELECT * from users WHERE api_key = $1"

	var user types.User

	err := p.DB.QueryRow(stmt, api_key).Scan(&user.ID, &user.Created_at, &user.Updated_at, &user.Name, &user.Api_key)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}
