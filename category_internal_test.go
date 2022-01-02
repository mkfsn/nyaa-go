package nyaa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory_validate(t *testing.T) {
	tests := []struct {
		p    Provider
		c    Category
		want error
	}{
		{ProviderNyaa, CategoryAll, nil},
		{ProviderNyaa, categoryNyaaBegin - 1, fmt.Errorf("invalid Category value for provider: ProviderNyaa")},
		{ProviderNyaa, categoryNyaaEnd + 1, fmt.Errorf("invalid Category value for provider: ProviderNyaa")},
		{ProviderSukebei, CategoryAll, nil},
		{ProviderSukebei, categorySukebeiBegin - 1, fmt.Errorf("invalid Category value for provider: ProviderSukebei")},
		{ProviderSukebei, categorySukebeiEnd + 1, fmt.Errorf("invalid Category value for provider: ProviderSukebei")},
		{providerEnd, CategoryAll, ErrUnknownProvider},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Provider %s Category %s", tt.p, tt.c)
		t.Run(name, func(t *testing.T) {
			got := tt.c.validate(tt.p)
			if tt.want != nil {
				assert.EqualError(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
