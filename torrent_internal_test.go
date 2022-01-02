package nyaa

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func Test_newTorrentFromDOM(t *testing.T) {
	type args struct {
		selection *goquery.Selection
		provider  Provider
	}

	type wants struct {
		torrent *Torrent
		err     error
	}

	type test struct {
		args  args
		wants wants
	}

	tests := map[string]func(*testing.T) test{
		"page with valid data": func(t *testing.T) test {
			html := `
				<table>
					<tr class="default">
						<td>
							<a href="/?c=4_3" title="Live Action - Non-English-translated">
								<img src="/static/img/icons/nyaa/4_3.png" alt="Live Action - Non-English-translated" class="category-icon">
							</a>
						</td>
						<td colspan="2">
							<a href="/view/1421189#comments" class="comments" title="1 comment">
								<i class="fa fa-comments-o"></i>1</a>
							<a href="/view/1421189" title="[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV">[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV</a>
						</td>
						<td class="text-center">
							<a href="/download/1421189.torrent"><i class="fa fa-fw fa-download"></i></a>
							<a href="magnet:?xt=urn:btih:ad81e1132f546b2643214e638e1fb9445b18597d&amp;dn=%5B%E6%B0%B4%E6%A8%B9%E5%A5%88%E3%80%85%5D%20Nana%20Mizuki%20-%20NANA%20CLIP%208%20BDMV&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
						</td>
						<td class="text-center">35.9 GiB</td>
						<td class="text-center" data-timestamp="1628985965" title="4 months 2 weeks 6 days 9 hours 18 minutes 12 seconds ago">2021-08-15 08:06</td>
						<td class="text-center">0</td>
						<td class="text-center">0</td>
						<td class="text-center">53</td>
					</tr>
				<table>
			`
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args: args{
					selection: doc.Selection,
					provider:  ProviderNyaa,
				},
				wants: wants{
					torrent: &Torrent{
						Category:           "Live Action - Non-English-translated",
						Name:               "[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV",
						Link:               "https://nyaa.si/download/1421189.torrent",
						Magnet:             "magnet:?xt=urn:btih:ad81e1132f546b2643214e638e1fb9445b18597d&dn=%5B%E6%B0%B4%E6%A8%B9%E5%A5%88%E3%80%85%5D%20Nana%20Mizuki%20-%20NANA%20CLIP%208%20BDMV&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce",
						Size:               "35.9 GiB",
						Date:               time.Unix(1628985965, 0),
						Seeders:            0,
						Leechers:           0,
						CompletedDownloads: 53,
						detailLink:         "https://nyaa.si/view/1421189",
					},
				},
			}
		},

		"empty data-timestamp": func(t *testing.T) test {
			html := `
				<table>
					<tr class="default">
						<td>
							<a href="/?c=4_3" title="Live Action - Non-English-translated">
								<img src="/static/img/icons/nyaa/4_3.png" alt="Live Action - Non-English-translated" class="category-icon">
							</a>
						</td>
						<td colspan="2">
							<a href="/view/1421189#comments" class="comments" title="1 comment">
								<i class="fa fa-comments-o"></i>1</a>
							<a href="/view/1421189" title="[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV">[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV</a>
						</td>
						<td class="text-center">
							<a href="/download/1421189.torrent"><i class="fa fa-fw fa-download"></i></a>
							<a href="magnet:?xt=urn:btih:ad81e1132f546b2643214e638e1fb9445b18597d&amp;dn=%5B%E6%B0%B4%E6%A8%B9%E5%A5%88%E3%80%85%5D%20Nana%20Mizuki%20-%20NANA%20CLIP%208%20BDMV&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
						</td>
						<td class="text-center">35.9 GiB</td>
						<td class="text-center" data-timestamp="" title="4 months 2 weeks 6 days 9 hours 18 minutes 12 seconds ago">2021-08-15 08:06</td>
						<td class="text-center">0</td>
						<td class="text-center">0</td>
						<td class="text-center">53</td>
					</tr>
				<table>
			`
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args: args{
					selection: doc.Selection,
					provider:  ProviderNyaa,
				},
				wants: wants{err: errors.New(`failed to parse data-timestamp: strconv.ParseInt: parsing "": invalid syntax`)},
			}
		},

		"empty seeders": func(t *testing.T) test {
			html := `
				<table>
					<tr class="default">
						<td>
							<a href="/?c=4_3" title="Live Action - Non-English-translated">
								<img src="/static/img/icons/nyaa/4_3.png" alt="Live Action - Non-English-translated" class="category-icon">
							</a>
						</td>
						<td colspan="2">
							<a href="/view/1421189#comments" class="comments" title="1 comment">
								<i class="fa fa-comments-o"></i>1</a>
							<a href="/view/1421189" title="[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV">[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV</a>
						</td>
						<td class="text-center">
							<a href="/download/1421189.torrent"><i class="fa fa-fw fa-download"></i></a>
							<a href="magnet:?xt=urn:btih:ad81e1132f546b2643214e638e1fb9445b18597d&amp;dn=%5B%E6%B0%B4%E6%A8%B9%E5%A5%88%E3%80%85%5D%20Nana%20Mizuki%20-%20NANA%20CLIP%208%20BDMV&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
						</td>
						<td class="text-center">35.9 GiB</td>
						<td class="text-center" data-timestamp="1628985965" title="4 months 2 weeks 6 days 9 hours 18 minutes 12 seconds ago">2021-08-15 08:06</td>
						<td class="text-center"></td>
						<td class="text-center">0</td>
						<td class="text-center">53</td>
					</tr>
				<table>
			`
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args: args{
					selection: doc.Selection,
					provider:  ProviderNyaa,
				},
				wants: wants{err: errors.New(`failed to parse seeders: strconv.ParseInt: parsing "": invalid syntax`)},
			}
		},

		"empty leechers": func(t *testing.T) test {
			html := `
				<table>
					<tr class="default">
						<td>
							<a href="/?c=4_3" title="Live Action - Non-English-translated">
								<img src="/static/img/icons/nyaa/4_3.png" alt="Live Action - Non-English-translated" class="category-icon">
							</a>
						</td>
						<td colspan="2">
							<a href="/view/1421189#comments" class="comments" title="1 comment">
								<i class="fa fa-comments-o"></i>1</a>
							<a href="/view/1421189" title="[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV">[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV</a>
						</td>
						<td class="text-center">
							<a href="/download/1421189.torrent"><i class="fa fa-fw fa-download"></i></a>
							<a href="magnet:?xt=urn:btih:ad81e1132f546b2643214e638e1fb9445b18597d&amp;dn=%5B%E6%B0%B4%E6%A8%B9%E5%A5%88%E3%80%85%5D%20Nana%20Mizuki%20-%20NANA%20CLIP%208%20BDMV&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
						</td>
						<td class="text-center">35.9 GiB</td>
						<td class="text-center" data-timestamp="1628985965" title="4 months 2 weeks 6 days 9 hours 18 minutes 12 seconds ago">2021-08-15 08:06</td>
						<td class="text-center">0</td>
						<td class="text-center"></td>
						<td class="text-center">53</td>
					</tr>
				<table>
			`
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args: args{
					selection: doc.Selection,
					provider:  ProviderNyaa,
				},
				wants: wants{err: errors.New(`failed to parse leechers: strconv.ParseInt: parsing "": invalid syntax`)},
			}
		},

		"empty completedDownloads": func(t *testing.T) test {
			html := `
				<table>
					<tr class="default">
						<td>
							<a href="/?c=4_3" title="Live Action - Non-English-translated">
								<img src="/static/img/icons/nyaa/4_3.png" alt="Live Action - Non-English-translated" class="category-icon">
							</a>
						</td>
						<td colspan="2">
							<a href="/view/1421189#comments" class="comments" title="1 comment">
								<i class="fa fa-comments-o"></i>1</a>
							<a href="/view/1421189" title="[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV">[水樹奈々] Nana Mizuki - NANA CLIP 8 BDMV</a>
						</td>
						<td class="text-center">
							<a href="/download/1421189.torrent"><i class="fa fa-fw fa-download"></i></a>
							<a href="magnet:?xt=urn:btih:ad81e1132f546b2643214e638e1fb9445b18597d&amp;dn=%5B%E6%B0%B4%E6%A8%B9%E5%A5%88%E3%80%85%5D%20Nana%20Mizuki%20-%20NANA%20CLIP%208%20BDMV&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
						</td>
						<td class="text-center">35.9 GiB</td>
						<td class="text-center" data-timestamp="1628985965" title="4 months 2 weeks 6 days 9 hours 18 minutes 12 seconds ago">2021-08-15 08:06</td>
						<td class="text-center">0</td>
						<td class="text-center">0</td>
						<td class="text-center"></td>
					</tr>
				<table>
			`
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			assert.NoError(t, err)

			return test{
				args: args{
					selection: doc.Selection,
					provider:  ProviderNyaa,
				},
				wants: wants{err: errors.New(`failed to parse completed downloads: strconv.ParseInt: parsing "": invalid syntax`)},
			}
		},
	}

	for name, fn := range tests {
		tt := fn(t)
		t.Run(name, func(t *testing.T) {
			got, err := newTorrentFromDOM(tt.args.selection, tt.args.provider)
			if tt.wants.err != nil {
				assert.EqualError(t, err, tt.wants.err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wants.torrent, got)
		})
	}
}
