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

	// Output:
	// &{Category:Live Action - Non-English-translated Name:[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV Link:https://nyaa.si/download/1421189.torrent Magnet:magnet:?xt=urn:btih:ad81e1132f546b2643214e638e1fb9445b18597d&dn=%5B%E6%B0%B4%E6%A8%B9%E5%A5%88%E3%80%85%5D%20Nana%20Mizuki%20-%20NANA%20CLIP%208%20BDMV&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce Size:35.9 GiB Date:2021-08-15 08:06:05 +0800 CST Seeders:0 Leechers:0 CompletedDownloads:53 detailLink:https://nyaa.si/view/1421189}
}
