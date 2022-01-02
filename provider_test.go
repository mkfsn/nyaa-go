package nyaa_test

import (
	"fmt"
	"testing"

	"github.com/mkfsn/nyaa-go"
)

func TestProvider_Host(t *testing.T) {
	tests := []struct {
		p    nyaa.Provider
		want string
	}{
		{nyaa.ProviderNyaa, "nyaa.si"},
		{nyaa.ProviderSukebei, "sukebei.nyaa.si"},
		{nyaa.Provider(-1), ""},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Provider %s", tt.p)
		t.Run(name, func(t *testing.T) {
			if got := tt.p.Host(); got != tt.want {
				t.Errorf("Host() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleProvider_Host() {
	fmt.Printf("%s\n", nyaa.ProviderNyaa.Host())
	fmt.Printf("%s\n", nyaa.ProviderSukebei.Host())

	// Output:
	// nyaa.si
	// sukebei.nyaa.si
}
