package nyaa_test

import (
	"fmt"
	"testing"

	"github.com/mkfsn/nyaa-go"
)

func TestFilterBy_Value(t *testing.T) {
	tests := []struct {
		f    nyaa.FilterBy
		want string
	}{
		{nyaa.FilterByNoFilter, "0"},
		{nyaa.FilterByNoRemakes, "1"},
		{nyaa.FilterByTrustedOnly, "2"},
		{nyaa.FilterBy(-1), ""},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("FilterBy %s", tt.f)
		t.Run(name, func(t *testing.T) {
			if got := tt.f.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
