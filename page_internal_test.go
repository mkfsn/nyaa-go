package nyaa

import (
	"errors"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func Test_newPageInfoFromDocument(t *testing.T) {
	type args struct {
		doc *goquery.Document
	}

	type wants struct {
		pageInfo *PageInfo
		err      error
	}

	type test struct {
		args  args
		wants wants
	}

	tests := map[string]func(t *testing.T) test{
		"document with 6 pages": func(t *testing.T) test {
			html := `
				<ul class="pagination">
					<li class="previous disabled unavailable"><a> « </a></li>
					<li class="active">
						<a href="#">1 <span class="sr-only">(current)</span></a>
					</li>
					<li><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=2">2</a></li>
					<li><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=3">3</a></li>
					<li><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=4">4</a></li>
					<li><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=5">5</a></li>
					<li><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=6">6</a></li>
					<li class="next"><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=2">»</a></li>
				</ul>
			`

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args:  args{doc: doc},
				wants: wants{pageInfo: &PageInfo{CurrentPage: 1, LastPage: 6}},
			}
		},

		"invalid current page": func(t *testing.T) test {
			html := `
				<ul class="pagination">
					<li class="previous disabled unavailable"><a> « </a></li>
					<li class="active"><a>invalid page number</a></li>
					<li><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=2">2</a></li>
					<li class="next"><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=2">»</a></li>
				</ul>
			`

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args:  args{doc: doc},
				wants: wants{err: errors.New(`strconv.Atoi: parsing "invalid page number": invalid syntax`)},
			}
		},

		"invalid last page": func(t *testing.T) test {
			html := `
				<ul class="pagination">
					<li class="previous disabled unavailable"><a> « </a></li>
					<li class="active"><a>1</a></li>
					<li><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=2">invalid page number</a></li>
					<li class="next"><a href="/?f=0&amp;c=0_0&amp;q=Nana+Mizuki+&amp;p=2">»</a></li>
				</ul>
			`

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args:  args{doc: doc},
				wants: wants{err: errors.New(`strconv.Atoi: parsing "invalid page number": invalid syntax`)},
			}
		},
	}
	for name, fn := range tests {
		tt := fn(t)
		t.Run(name, func(t *testing.T) {
			got, err := newPageInfoFromDocument(tt.args.doc)
			if tt.wants.err != nil {
				assert.EqualError(t, err, tt.wants.err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, got, tt.wants.pageInfo)
		})
	}
}
