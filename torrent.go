package nyaa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Torrent represents a torrent with information retrieved from Nyaa.
type Torrent struct {
	Category           string    `json:"category"`
	Name               string    `json:"name"`
	Link               string    `json:"link"`
	Magnet             string    `json:"magnet"`
	Size               string    `json:"size"`
	Date               time.Time `json:"date"`
	Seeders            int64     `json:"seeders"`
	Leechers           int64     `json:"leechers"`
	CompletedDownloads int64     `json:"completeDownloads"`

	detailLink string
}

func newTorrentFromDOM(selection *goquery.Selection, provider Provider) (*Torrent, error) {
	td := selection.Find("td")

	category := td.Eq(0).Find("a").AttrOr("title", "")

	name := td.Eq(1).Find("a").Not(".comments").AttrOr("title", "")

	detailLink := fmt.Sprintf("%s%s", provider.BaseURL().String(), td.Eq(1).Find("a").Not(".comments").AttrOr("href", ""))

	link := fmt.Sprintf("%s%s", provider.BaseURL().String(), td.Eq(2).Find("a").Eq(0).AttrOr("href", ""))

	magnet := td.Eq(2).Find("a").Eq(1).AttrOr("href", "")

	size := td.Eq(3).Text()

	dataTimestamp, err := strconv.ParseInt(td.Eq(4).AttrOr("data-timestamp", ""), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse data-timestamp: %w", err)
	}

	date := time.Unix(dataTimestamp, 0)

	seeders, err := strconv.ParseInt(td.Eq(5).Text(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse seeders: %w", err)
	}

	leechers, err := strconv.ParseInt(td.Eq(6).Text(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse leechers: %w", err)
	}

	completedDownloads, err := strconv.ParseInt(td.Eq(7).Text(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse completed downloads: %w", err)
	}

	return &Torrent{
		Category:           category,
		Name:               name,
		Link:               link,
		Magnet:             magnet,
		Size:               size,
		Date:               date,
		Seeders:            seeders,
		Leechers:           leechers,
		CompletedDownloads: completedDownloads,
		detailLink:         detailLink,
	}, nil
}

// func (t *Torrent) GetDetails() (*TorrentDetails, error) {
// 	panic("not implemented")
// }

// type TorrentDetails struct {
// Submitter string
// Information string
// Comments []Comment
// Description string
// InfoHash string
// }
