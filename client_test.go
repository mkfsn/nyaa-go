package nyaa_test

import (
	"context"
	"fmt"
	"log"

	"github.com/mkfsn/nyaa-go"
)

func ExampleClient_Search() {
	c := nyaa.NewClient()
	torrents, _, err := c.Search(context.Background(), nyaa.SearchOptions{
		Provider:  nyaa.ProviderNyaa,
		FilterBy:  nyaa.FilterByNoFilter,
		Category:  nyaa.CategoryAll,
		Query:     "Nana Mizuki - NANA CLIP 8 BDMV",
		SortBy:    nyaa.SortByDate,
		SortOrder: nyaa.SortOrderDesc,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for _, torrent := range torrents {
		fmt.Printf("%s\n", torrent.Name)
		fmt.Printf(" - %s\n", torrent.Category)
		fmt.Printf(" - %s\n", torrent.Size)
		fmt.Printf(" - %s\n", torrent.Link)
		fmt.Printf(" - %s\n", torrent.Date.UTC())
		// fmt.Printf(" - %d\n", torrent.Seeders)
		// fmt.Printf(" - %d\n", torrent.Leechers)
		// fmt.Printf(" - %d\n", torrent.CompletedDownloads)
	}

	// Output:
	// [水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV
	//  - Live Action - Non-English-translated
	//  - 35.9 GiB
	//  - https://nyaa.si/download/1421189.torrent
	//  - 2021-08-15 00:06:05 +0000 UTC
}
