package scrapper

import (
	"encoding/xml"
	"github.com/ante-neh/Rss-aggregator/types"
	"io"
	"net/http"
	"time"
)

func UrlToFeed(url string) (types.RssFeed, error) {

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := httpClient.Get(url)

	if err != nil {
		return types.RssFeed{}, nil
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		return types.RssFeed{}, nil
	}

	xmlResponse := types.RssFeed{}

	err = xml.Unmarshal(data, &xmlResponse)

	if err != nil {
		return types.RssFeed{}, nil
	}

	return xmlResponse, nil

}
