package nyaa

import (
	"errors"
	"net/url"
)

// Provider represents the supported sites of Nyaa.
type Provider int

// Currently supported providers:
//  - ProviderNyaa: http://nyaa.si
//  - ProviderSukebei: http://sukebei.nyaa.si (NSFW)
const (
	ProviderNyaa Provider = iota
	ProviderSukebei
	providerEnd
)

// String implements fmt.Stringer interface.
func (p Provider) String() string {
	switch p {
	case ProviderNyaa:
		return "ProviderNyaa"
	case ProviderSukebei:
		return "ProviderSukebei"
	}

	return ""
}

// BaseURL returns the base URL of the provider.
func (p Provider) BaseURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   p.Host(),
	}
}

// Host returns the host of the provider.
func (p Provider) Host() string {
	switch p {
	case ProviderNyaa:
		return "nyaa.si"
	case ProviderSukebei:
		return "sukebei.nyaa.si"
	}

	return ""
}

func (p Provider) validate() error {
	if p >= providerEnd {
		return errors.New("invalid Provider value")
	}

	return nil
}
