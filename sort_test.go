package nyaa_test

import (
	"fmt"
	"testing"

	"github.com/mkfsn/nyaa-go"
)

func TestSortBy_Value(t *testing.T) {
	tests := []struct {
		s    nyaa.SortBy
		want string
	}{
		{nyaa.SortByDate, "id"},
		{nyaa.SortByComments, "comments"},
		{nyaa.SortByDownloads, "downloads"},
		{nyaa.SortBySeeders, "seeders"},
		{nyaa.SortByLeechers, "leechers"},
		{nyaa.SortBySize, "size"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("SortBy %s", tt.s)
		t.Run(name, func(t *testing.T) {
			if got := tt.s.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortOrder_Value(t *testing.T) {
	tests := []struct {
		s    nyaa.SortOrder
		want string
	}{
		{nyaa.SortOrderDesc, "desc"},
		{nyaa.SortOrderAsc, "asc"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("SortOrder %s", tt.s)
		t.Run(name, func(t *testing.T) {
			if got := tt.s.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
