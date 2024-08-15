package database

import (
	"database/sql"
	"time"
	"github.com/google/uuid"
)

type DatabaseOperation interface{
	Createuser(id uuid.UUID, created_at time.Time, updated_at time.Time, name string) (int, error)
}
type Postgres struct {
	DB *sql.DB
}


func (p *Postgres) Createuser(id uuid.UUID, created_at time.Time, updated_at time.Time, name string) (int, error){
	stmt := "INSERT INTO users(id, created_at, updated_at, name) VALUES($1, $2, $3, $4)"
	result, err := p.DB.Exec(stmt, id, created_at, updated_at, name)

	if err != nil{
		return 0, err
	}

	lastId, err := result.LastInsertId() 
	if err != nil{
		return 0, err
	}

	return int(lastId), nil 
}
