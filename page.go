package nyaa

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Page represents the page number in both request and response.
type Page int

// Value returns the value of the query parameter in the HTTP request.
func (p Page) Value() string {
	return strconv.Itoa(int(p))
}

// PageInfo represents the pagination information parsed from the web page.
type PageInfo struct {
	IsSinglePage bool
	CurrentPage  Page
	LastPage     Page
}

func newPageInfoFromDocument(doc *goquery.Document) (*PageInfo, error) {
	pagination := doc.Find("ul.pagination")

	if pagination.Length() == 0 {
		return &PageInfo{IsSinglePage: true}, nil
	}

	currentPage, err := strconv.Atoi(strings.TrimSpace(pagination.Find("li.active > a").Children().Remove().End().Text()))
	if err != nil {
		return nil, err
	}

	lastPage, err := strconv.Atoi(pagination.Find("li").Not(".next").Last().Text())
	if err != nil {
		return nil, err
	}

	return &PageInfo{
		CurrentPage: Page(currentPage),
		LastPage:    Page(lastPage),
	}, nil
}
