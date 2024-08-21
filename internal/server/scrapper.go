package server

import (
	"strings"
	"sync"
	"time"

	"github.com/ante-neh/Rss-aggregator/internal/scrapper"
	"github.com/ante-neh/Rss-aggregator/types"
	"github.com/google/uuid"
)

func (s *Server) StartScrapping(concurrency int, timeBetweenRequest time.Duration){
	timeTicker := time.NewTicker(timeBetweenRequest) 

	for ;; <-timeTicker.C{
		feeds, err := s.DB.GetNextFeedToFetch(concurrency)
		if err != nil{
			s.ErrorLogger.Println(err)
			continue 
		}

		wg := &sync.WaitGroup{}
		
		for _, feed := range feeds{
			wg.Add(1)
			go s.ScrapeFeed(wg, *feed) 
		}

		wg.Wait()

	}
}


func (s *Server) ScrapeFeed(wg *sync.WaitGroup, feed types.Feeds){
	defer wg.Done() 
	_, err := s.DB.MarkFeedsAsFetch(feed.ID) 
	if err != nil{
		s.ErrorLogger.Println(err)
		return 
	}

	rssFeed, err := scrapper.UrlToFeed(feed.Url)

	if err != nil{
		s.ErrorLogger.Println(err)
		return 
	}

	for _, item := range rssFeed.Channel.Item{
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate) 
		
		if err != nil{
			s.ErrorLogger.Println(err)
			continue 
		}

		err = s.DB.CreatePost(uuid.New(), feed.ID, time.Now().UTC(), time.Now().UTC(), pubAt, item.Title, item.Description, item.Link)

		if err != nil{
			if strings.Contains(err.Error(), "duplicate key"){
				continue 
			}

			s.ErrorLogger.Println(err)
			continue 
		}
	}
}