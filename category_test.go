package nyaa_test

import (
	"fmt"
	"testing"

	"github.com/mkfsn/nyaa-go"
)

func TestCategory_Value(t *testing.T) {
	tests := []struct {
		p    nyaa.Provider
		c    nyaa.Category
		want string
	}{
		// For ProviderNyaa
		{nyaa.ProviderNyaa, nyaa.CategoryAll, "0_0"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAnime, "1_0"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAnimeMusicVideo, "1_1"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAnimeEnglishTranslated, "1_2"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAnimeNonEnglishTranslated, "1_3"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAnimeRaw, "1_4"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAudio, "2_0"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAudioLossless, "2_1"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaAudioLossy, "2_2"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaLiterature, "3_0"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaLiteratureEnglishTranslated, "3_1"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaLiteratureNonEnglishTranslated, "3_2"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaLiteratureRaw, "3_3"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaLiveAction, "4_0"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaLiveActionEnglishTranslated, "4_1"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaPictures, "5_0"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaPicturesGraphics, "5_1"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaPicturesPhotos, "5_2"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaSoftware, "6_0"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaSoftwareApplications, "6_1"},
		{nyaa.ProviderNyaa, nyaa.CategoryNyaaSoftwareGames, "6_2"},
		// For ProviderSukebei
		{nyaa.ProviderSukebei, nyaa.CategoryAll, "0_0"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiArt, "1_0"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiArtAnime, "1_1"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiArtDoujinshi, "1_2"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiArtGames, "1_3"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiArtManga, "1_4"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiArtPictures, "1_5"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiRealLife, "2_0"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiRealLifePhotobooksAndPictures, "2_1"},
		{nyaa.ProviderSukebei, nyaa.CategorySukebeiRealLifeVideos, "2_2"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Provider %s Category %s", tt.p, tt.c)
		t.Run(name, func(t *testing.T) {
			if got := tt.c.Value(tt.p); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
