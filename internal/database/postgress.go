package database

import (
	"database/sql"
	"time"
	"github.com/ante-neh/Rss-aggregator/types"
	"github.com/google/uuid"
)

type DatabaseOperation interface{
	Createuser(id uuid.UUID, created_at time.Time, updated_at time.Time, name string) (types.User, error)
}
type Postgres struct {
	DB *sql.DB
}


func (p *Postgres) Createuser(id uuid.UUID, created_at time.Time, updated_at time.Time, name string) (types.User, error){
	stmt := "INSERT INTO users(id, created_at, updated_at, name) VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at, name"

	var user types.User 
	err := p.DB.QueryRow(stmt, id, created_at, updated_at, name).Scan(&user.ID, &user.Created_at, &user.Updated_at, &user.Name )

	if err != nil{
		return types.User{}, err
	}

	return user, nil 
}
