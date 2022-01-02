package nyaa

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchOptions_validate(t *testing.T) {
	tests := []struct {
		s    SearchOptions
		want error
	}{
		{SearchOptions{Provider: Provider(-1)}, ErrUnknownProvider},
		{SearchOptions{FilterBy: FilterBy(-1)}, ErrUnknownFilterBy},
		{SearchOptions{Provider: ProviderNyaa, Category: Category(-1)}, errors.New("invalid Category value for provider: ProviderNyaa")},
		{SearchOptions{Provider: ProviderSukebei, Category: Category(-1)}, errors.New("invalid Category value for provider: ProviderSukebei")},
		{SearchOptions{SortBy: SortBy(-1)}, ErrUnknownSortBy},
		{SearchOptions{SortOrder: SortOrder(-1)}, ErrUnknownSortOrder},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("SearchOptions %v", tt.s)
		t.Run(name, func(t *testing.T) {
			got := tt.s.validate()
			if tt.want != nil {
				assert.EqualError(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
