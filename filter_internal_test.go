package nyaa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterBy_validate(t *testing.T) {
	tests := []struct {
		f    FilterBy
		want error
	}{
		{FilterByNoFilter, nil},
		{FilterByNoRemakes, nil},
		{FilterByTrustedOnly, nil},
		{filterByEnd, ErrUnknownFilterBy},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("FilterBy %s", tt.f)
		t.Run(name, func(t *testing.T) {
			got := tt.f.validate()
			if tt.want != nil {
				assert.EqualError(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
