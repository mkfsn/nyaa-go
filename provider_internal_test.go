package nyaa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider_validate(t *testing.T) {
	tests := []struct {
		p    Provider
		want error
	}{
		{ProviderNyaa, nil},
		{ProviderSukebei, nil},
		{providerEnd, ErrUnknownProvider},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Provider %s", tt.p)
		t.Run(name, func(t *testing.T) {
			got := tt.p.validate()
			if tt.want != nil {
				assert.EqualError(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
