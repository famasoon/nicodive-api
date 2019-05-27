package api

import (
	"log"

	"github.com/mmcdole/gofeed"
)

// Endpoint that gets the ranking of Nico Nico Douga in view count order
// e.g. http://www.nicovideo.jp/ranking/view/hourly/game?rss=2.0
const VIEW_RANKING_ENDPOINT = "http://www.nicovideo.jp/ranking/view/"

type RankingItem struct {
	Title       string `json:"title"`
	Rank        int    `json:"rank"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"desc"`
	Link        string `json:"link"`
}

func GetRanking(category, period string) ([]RankingItem, error) {
	rankingURL := VIEW_RANKING_ENDPOINT + "/" + period + "/" + category + "?rss=2.0"

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rankingURL)
	if err != nil {
		log.Printf("[ERROR] failed to parse ranking, rankingURL: %s\n", rankingURL)
		return nil, err
	}

	itemChan := make(chan RankingItem)
	defer close(itemChan)

	errChan := make(chan error, 1)
	defer close(errChan)

	items := make([]RankingItem, len(feed.Items))

	for i := range feed.Items {
		go func(link string, rank int) {
			videoID := GetVideoID(link)
			videoInfo, err := GetVideoInfo(videoID)
			if err != nil {
				errChan <- err
			}

			itemChan <- RankingItem{
				Title:       videoInfo.Video.Title,
				Rank:        rank + 1,
				Thumbnail:   videoInfo.Video.ThumbnailURL,
				Description: videoInfo.Video.Description,
				Link:        videoID,
			}
		}(feed.Items[i].Link, i)
	}

	for range feed.Items {
		select {
		case item := <-itemChan:
			items[item.Rank-1] = item
		case err := <-errChan:
			log.Printf("[ERROR] failed to get ranking items, rankingURL: %s\n", rankingURL)
			return nil, err
		}
	}
	return items, nil
}
