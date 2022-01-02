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
	categoryNyaaBegin                          Category = 0x100
	CategoryNyaaAnime                          Category = 0x110
	CategoryNyaaAnimeMusicVideo                Category = 0x111
	CategoryNyaaAnimeEnglishTranslated         Category = 0x112
	CategoryNyaaAnimeNonEnglishTranslated      Category = 0x113
	CategoryNyaaAnimeRaw                       Category = 0x114
	CategoryNyaaAudio                          Category = 0x120
	CategoryNyaaAudioLossless                  Category = 0x121
	CategoryNyaaAudioLossy                     Category = 0x122
	CategoryNyaaLiterature                     Category = 0x130
	CategoryNyaaLiteratureEnglishTranslated    Category = 0x131
	CategoryNyaaLiteratureNonEnglishTranslated Category = 0x132
	CategoryNyaaLiteratureRaw                  Category = 0x133
	CategoryNyaaLiveAction                     Category = 0x140
	CategoryNyaaLiveActionEnglishTranslated    Category = 0x141
	CategoryNyaaPictures                       Category = 0x150
	CategoryNyaaPicturesGraphics               Category = 0x151
	CategoryNyaaPicturesPhotos                 Category = 0x152
	CategoryNyaaSoftware                       Category = 0x160
	CategoryNyaaSoftwareApplications           Category = 0x161
	CategoryNyaaSoftwareGames                  Category = 0x162
	categoryNyaaEnd                            Category = 0x1FF
	// Categories between categorySukebeiBegin and categorySukebeiEnd are for ProviderSukebei.
	categorySukebeiBegin                         Category = 0x200
	CategorySukebeiArt                           Category = 0x210
	CategorySukebeiArtAnime                      Category = 0x211
	CategorySukebeiArtDoujinshi                  Category = 0x212
	CategorySukebeiArtGames                      Category = 0x213
	CategorySukebeiArtManga                      Category = 0x214
	CategorySukebeiArtPictures                   Category = 0x215
	CategorySukebeiRealLife                      Category = 0x220
	CategorySukebeiRealLifePhotobooksAndPictures Category = 0x221
	CategorySukebeiRealLifeVideos                Category = 0x222
	categorySukebeiEnd                           Category = 0x2FF
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
