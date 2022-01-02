package nyaa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortBy_validate(t *testing.T) {
	tests := []struct {
		s    SortBy
		want error
	}{
		{SortByDate, nil},
		{sortByEnd, ErrUnknownSortBy},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("SortBy %s", tt.s)
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

func TestSortOrder_validate(t *testing.T) {
	tests := []struct {
		s    SortOrder
		want error
	}{
		{SortOrderDesc, nil},
		{sortOrderEnd, ErrUnknownSortOrder},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("SortOrder %s", tt.s)
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
