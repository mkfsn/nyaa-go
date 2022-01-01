# nyaa-go

[![Go Report Card](https://goreportcard.com/badge/github.com/mkfsn/nyaa-go)](https://goreportcard.com/report/github.com/mkfsn/nyaa-go)
[![Actions Status](https://github.com/mkfsn/nyaa-go/actions/workflows/develop.yaml/badge.svg)](https://github.com/mkfsn/nyaa-go/actions)
[![codecov](https://codecov.io/gh/mkfsn/nyaa-go/branch/develop/graph/badge.svg?token=Z3IKJYGSJV)](https://codecov.io/gh/mkfsn/nyaa-go)

[![Go Reference](https://pkg.go.dev/badge/github.com/mkfsn/nyaa-go.svg)](https://pkg.go.dev/github.com/mkfsn/nyaa-go)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mkfsn/nyaa-go)
[![License](https://img.shields.io/github/license/mkfsn/notion-go.svg)](./LICENSE.md)




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
