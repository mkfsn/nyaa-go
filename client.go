package nyaa

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Client represents a Nyaa client.
type Client struct {
	httpClient *http.Client
}

// NewClient returns a Nyaa client.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// Search sends a request based on the given SearchOptions and returns the matched torrents information.
func (c *Client) Search(ctx context.Context, opts SearchOptions) ([]*Torrent, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, opts.buildURL().String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	rows := doc.Find("table.torrent-list > tbody > tr")

	torrents := make([]*Torrent, 0, rows.Length())

	rows.Each(func(i int, selection *goquery.Selection) {
		torrent, err := newTorrentFromDOM(selection, opts.Provider)
		if err != nil {
			return
		}

		torrents = append(torrents, torrent)
	})

	return torrents, nil
}
