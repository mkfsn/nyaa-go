package nyaa

import (
	"net/url"
)

// Provider represents the supported sites of Nyaa.
type Provider int

// Currently supported providers.
const (
	// ProviderNyaa is for http://nyaa.si.
	ProviderNyaa Provider = iota
	// ProviderSukebei is for http://sukebei.nyaa.si (NSFW).
	ProviderSukebei
	// providerEnd is for border check.
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

	return unknownEntityName
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
		return ErrUnknownProvider
	}

	return nil
}
