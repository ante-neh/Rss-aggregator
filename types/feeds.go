package types

import (
	"github.com/google/uuid"
	"time"
)

type Feeds struct {
	ID         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"user_id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Url        string    `json:"url"`
	Last_Fetched_At string `json:"last_fetched_at"`
}

type FeedFollow struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"creaeted_at"`
	Updated_at time.Time `json:"updated_at"`
	FeedId     uuid.UUID `json:"feed_id"`
	UserId     uuid.UUID `json:"user_id"`
}

type RssItem struct{
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubdate"`
}
type Channel struct{
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	Language string `xml:"language"`
	Item []RssItem `xml:"item"`

}

type RssFeed struct{
	Channel Channel `xml:"channel"`
}
