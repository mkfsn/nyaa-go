package nyaa_test

import (
	"fmt"
	"log"

	"github.com/mkfsn/nyaa-go"
)

func ExampleClient_Search() {
	c := nyaa.NewClient()
	torrents, err := c.Search(nyaa.SearchOptions{
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
		fmt.Printf("%+v\n", torrent)
	}
}
