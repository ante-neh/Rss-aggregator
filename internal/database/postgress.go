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
	CreateFeeds(id uuid.UUID, user_id uuid.UUID, created_at, update_at time.Time, name, url string) (types.Feeds, error)
	GetFeeds() ([]types.Feeds, error)
	CreateFeedFollows(id uuid.UUID, created_at, updated_at time.Time, feed_id, user_id uuid.UUID)(types.FeedFollow, error)
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

func (p *Postgres) CreateFeeds(id, user_id uuid.UUID, name, url string, created_at, updated_at time.Time) (types.Feeds, error) {
	stmt := "INSERT INTO feeds(id, created_at, updated_at, name, url, user_id) VALUES($1, $2, $3, $4, $5, $6) RETURNING *"
	var feed types.Feeds
	err := p.DB.QueryRow(stmt, id, created_at, updated_at, name, url, user_id).Scan(&feed.ID, &feed.Created_at, &feed.Updated_at, &feed.Name, &feed.Url, &feed.UserId)

	if err != nil {
		return types.Feeds{}, nil
	}

	return feed, nil
}

func (p *Postgres) GetFeeds() ([]*types.Feeds, error) {
	stmt := "SELECT * FROM feeds"
	rows, err := p.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feeds []*types.Feeds

	for rows.Next() {
		feed := &types.Feeds{}
		err := rows.Scan(&feed.ID, &feed.Created_at, &feed.Updated_at, &feed.Name, &feed.Url, &feed.UserId)

		if err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return feeds, nil
}


func (p *Postgres) CreateFeedFollows(id uuid.UUID, created_at, updated_at time.Time, feed_id, user_id uuid.UUID) (types.FeedFollow, error){
	return types.FeedFollow{}, nil 
}