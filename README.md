# nyaa-go

An unofficial nyaa.si client library for Go


## Installation

```bash
go get github.com/mkfsn/nyaa-go
```

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/mkfsn/nyaa-go"
)

func main() {
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
```
