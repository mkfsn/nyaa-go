package nyaa

import (
	"fmt"
)

// Category represents the category for searching torrents in Nyaa.
// The available category varies by different providers.
type Category int

// The categories for the currently supported providers.
const (
	CategoryAll Category = 0x000
	// Categories between categoryNyaaBegin and categoryNyaaEnd are for ProviderNyaa.
	categoryNyaaBegin                          = 0x100
	CategoryNyaaAnime                          = 0x110
	CategoryNyaaAnimeMusicVideo                = 0x111
	CategoryNyaaAnimeEnglishTranslated         = 0x112
	CategoryNyaaAnimeNonEnglishTranslated      = 0x113
	CategoryNyaaAnimeRaw                       = 0x114
	CategoryNyaaAudio                          = 0x120
	CategoryNyaaAudioLossless                  = 0x121
	CategoryNyaaAudioLossy                     = 0x122
	CategoryNyaaLiterature                     = 0x130
	CategoryNyaaLiteratureEnglishTranslated    = 0x131
	CategoryNyaaLiteratureNonEnglishTranslated = 0x132
	CategoryNyaaLiteratureRaw                  = 0x133
	CategoryNyaaLiveAction                     = 0x140
	CategoryNyaaLiveActionEnglishTranslated    = 0x141
	CategoryNyaaPictures                       = 0x150
	CategoryNyaaPicturesGraphics               = 0x151
	CategoryNyaaPicturesPhotos                 = 0x152
	CategoryNyaaSoftware                       = 0x160
	CategoryNyaaSoftwareApplications           = 0x161
	CategoryNyaaSoftwareGames                  = 0x162
	categoryNyaaEnd                            = 0x1FF
	// Categories between categorySukebeiBegin and categorySukebeiEnd are for ProviderSukebei.
	categorySukebeiBegin                         = 0x200
	CategorySukebeiArt                           = 0x210
	CategorySukebeiArtAnime                      = 0x211
	CategorySukebeiArtDoujinshi                  = 0x212
	CategorySukebeiArtGames                      = 0x213
	CategorySukebeiArtManga                      = 0x214
	CategorySukebeiArtPictures                   = 0x215
	CategorySukebeiRealLife                      = 0x220
	CategorySukebeiRealLifePhotobooksAndPictures = 0x221
	CategorySukebeiRealLifeVideos                = 0x222
	categorySukebeiEnd                           = 0x2FF
)

var (
	categoryNames = map[Category]string{
		CategoryAll:                                  "All",
		CategoryNyaaAnime:                            "NyaaAnimeAll",
		CategoryNyaaAnimeMusicVideo:                  "NyaaAnimeMusicVideo",
		CategoryNyaaAnimeEnglishTranslated:           "NyaaAnimeEnglishTranslated",
		CategoryNyaaAnimeNonEnglishTranslated:        "NyaaAnimeNonEnglishTranslated",
		CategoryNyaaAnimeRaw:                         "NyaaAnimeRaw",
		CategoryNyaaAudio:                            "NyaaAudioAll",
		CategoryNyaaAudioLossless:                    "NyaaAudioLossless",
		CategoryNyaaAudioLossy:                       "NyaaAudioLossy",
		CategoryNyaaLiterature:                       "NyaaLiteratureAll",
		CategoryNyaaLiteratureEnglishTranslated:      "NyaaLiteratureEnglishTranslated",
		CategoryNyaaLiteratureNonEnglishTranslated:   "NyaaLiteratureNonEnglishTranslated",
		CategoryNyaaLiteratureRaw:                    "NyaaLiteratureRaw",
		CategoryNyaaLiveAction:                       "NyaaLiveActionAll",
		CategoryNyaaLiveActionEnglishTranslated:      "NyaaLiveActionEnglishTranslated",
		CategoryNyaaPictures:                         "NyaaPicturesAll",
		CategoryNyaaPicturesGraphics:                 "NyaaPicturesGraphics",
		CategoryNyaaPicturesPhotos:                   "NyaaPicturesPhotos",
		CategoryNyaaSoftware:                         "NyaaSoftwareAll",
		CategoryNyaaSoftwareApplications:             "NyaaSoftwareApplications",
		CategoryNyaaSoftwareGames:                    "NyaaSoftwareGames",
		CategorySukebeiArt:                           "SukebeiArtAll",
		CategorySukebeiArtAnime:                      "SukebeiArtAnime",
		CategorySukebeiArtDoujinshi:                  "SukebeiArtDoujinshi",
		CategorySukebeiArtGames:                      "SukebeiArtGames",
		CategorySukebeiArtManga:                      "SukebeiArtManga",
		CategorySukebeiArtPictures:                   "SukebeiArtPictures",
		CategorySukebeiRealLife:                      "SukebeiRealLifeAll",
		CategorySukebeiRealLifePhotobooksAndPictures: "SukebeiRealLifePhotobooksAndPictures",
		CategorySukebeiRealLifeVideos:                "SukebeiRealLifeVideos",
	}
)

// Value returns the value of the query parameter in the HTTP request based on the given Provider.
func (c Category) Value(p Provider) string {
	if c == CategoryAll {
		return "0_0"
	}

	switch p {
	case ProviderNyaa:
		base := c - categoryNyaaBegin

		return fmt.Sprintf("%d_%d", base/16, base%16)

	case ProviderSukebei:
		base := c - categorySukebeiBegin

		return fmt.Sprintf("%d_%d", base/16, base%16)
	}

	return ""
}

// String implements fmt.Stringer interface.
func (c Category) String() string {
	if name, ok := categoryNames[c]; ok {
		return name
	}

	return ""
}

func (c Category) validate(p Provider) error {
	if c == CategoryAll {
		return nil
	}

	switch p {
	case ProviderNyaa:
		if c <= categoryNyaaBegin || c >= categoryNyaaEnd {
			return fmt.Errorf("invalid Category value for provider: %s", p)
		}

	case ProviderSukebei:
		if c <= categorySukebeiBegin || c >= categorySukebeiEnd {
			return fmt.Errorf("invalid Category value for provider: %s", p)
		}
	}

	return nil
}
