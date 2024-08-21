package database

import (
	"database/sql"
	"time"

	"github.com/ante-neh/Rss-aggregator/types"
	"github.com/ante-neh/Rss-aggregator/util"
	"github.com/google/uuid"
)

type DatabaseOperation interface {
	GetFeeds()([]types.Feeds, error)
	DeleteFeedFollow(id uuid.UUID)(error)
	GetUser(api_key string) (types.User, error)
	MarkFeedsAsFetch(id uuid.UUID)(types.Feeds, error)
	GetNextFeedsToFetch(limit int)([]types.Feeds, error)
	GetFeedFollows(id uuid.UUID)([]types.FeedFollow, error)
	Createuser(id uuid.UUID, created_at time.Time, updated_at time.Time, name string) (types.User, error)
	CreateFeeds(id uuid.UUID, user_id uuid.UUID, created_at, update_at time.Time, name, url string) (types.Feeds, error)
	CreateFeedFollows(id uuid.UUID, created_at, updated_at time.Time, feed_id, user_id uuid.UUID)(types.FeedFollow, error)
	CreatePost(id, feed_id uuid.UUID, created_at, updated_at, published_at time.Time, title, description, url string)(error)
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
	stmt := "INSERT INTO feed_follows(id, created_at, updated_at, feed_id, user_id) VALUES($1, $2, $3, $4, $5) RETURNING *"

	var feedFollow types.FeedFollow 
	err := p.DB.QueryRow(stmt, id, created_at, updated_at, feed_id, user_id).Scan(&feedFollow.ID, &feedFollow.Created_at, &feedFollow.Updated_at, &feedFollow.FeedId, &feedFollow.UserId)

	if err != nil{
		return types.FeedFollow{}, err

	}

	return feedFollow, nil 
}


func(p *Postgres) GetFeedFollows(id uuid.UUID)([]*types.FeedFollow, error){
	stmt :="SELECT * FROM feed_follows WHERE user_id = $1"

	rows, err := p.DB.Query(stmt, id) 
	if err != nil{
		return nil ,err
	}

	defer rows.Close() 

	var feedFollows []*types.FeedFollow  

	for rows.Next(){
		feedFollow := &types.FeedFollow{}
		err := rows.Scan(&feedFollow.ID, &feedFollow.Created_at, &feedFollow.Updated_at, &feedFollow.FeedId, &feedFollow.UserId)
		
		if err != nil{
			return nil, err
		}

		feedFollows = append(feedFollows, feedFollow)
	}

	if err := rows.Err(); err != nil{
		return nil, err
	}

	return feedFollows, nil
}


func(p *Postgres) DeleteFeedFollow(id uuid.UUID)(error){
	stmt := "DELETE FROM feed_follows WHERE user_id = $1"
	_, err := p.DB.Exec(stmt, id)
	
	if err != nil{
		return err
	}

	return nil 
}

func (p *Postgres) GetNextFeedToFetch(limit int)([]*types.Feeds, error){
	stmt := "SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1"

	rows, err := p.DB.Query(stmt, limit)
	if err != nil{
		return []*types.Feeds{}, err
	}
	defer rows.Close()

	feeds := []*types.Feeds{}

	for rows.Next(){
		feed := &types.Feeds{}
		err := rows.Scan(&feed.ID, &feed.Created_at, &feed.Updated_at, &feed.Name, &feed.Url,  &feed.UserId,  &feed.Last_Fetched_At)

		if err != nil{
			return []*types.Feeds{}, err
		}

		feeds = append(feeds, feed)
	}

	if err := rows.Err(); err != nil{

		return []*types.Feeds{}, err
	}


	return feeds, nil
}


func (p *Postgres) MarkFeedsAsFetch(id uuid.UUID)(types.Feeds, error){
	stmt := "UPDATE feeds SET last_fetched_at=$1, updated_at=$2 WHERE id=$3 RETURNING id, user_id, name, created_at, updated_at, url, last_fetched_at"
	var feed types.Feeds

	err := p.DB.QueryRow(stmt, time.Now(), time.Now(), id).Scan(
		&feed.ID,
		&feed.UserId,
		&feed.Name,
		&feed.Created_at,
		&feed.Updated_at,
		&feed.Url,
		&feed.Last_Fetched_At,
	)
	if err != nil {
		return types.Feeds{}, err
	} 

	return feed, nil 
}


func (p *Postgres) CreatePost(id, feeds_id uuid.UUID, created_at, updated_at, published_at time.Time, title, description, url string)(error){

	stmt := "INSERT INTO posts(id, created_at, updated_at, title, description, published_at, url, feeds_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := p.DB.Query(stmt, id, created_at, updated_at, title, description, published_at, url, feeds_id)
	if err != nil{
		return err 

	}

	return nil 
}


func (p *Postgres) GetFeedsForUser(id uuid.UUID, limit int)([]*types.Post, error){
	stmt := "SELECT posts.* FROM posts JOIN posts.feeds_id == feed_follows.feed_id WHERE feed_follows.user_id = $1 ORDER BY posts.published_at DESC LIMIT = $2"

	rows, err := p.DB.Query(stmt, id, limit)

	if err != nil{
		return nil, err 
	}

	defer rows.Close() 
	posts := []*types.Post{}

	for rows.Next(){
		post := &types.Post{}
		err := rows.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Title, &post.Url,  &post.PublishedAt,  &post.Description, &post.FeedId)

		if err != nil{
			return []*types.Post{}, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil{

		return []*types.Post{}, err
	}


	return posts, nil
}