package database

import (
	"database/sql"
)

type DatabaseOperation interface{
	createuser(id, created_at, updated_at, name string) (int, error)
}
type Postgres struct {
	DB *sql.DB
}


func (p *Postgres) createuser(id, created_at, updated_at, name string) (int, error){
	stmt := "INSERT INTO users(id, creaeted_at, updated_at, name) VALUES($1, $2, $3, $4)"
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
