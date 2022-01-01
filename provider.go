package nyaa

import (
	"errors"
	"net/url"
)

// Provider represents the supported sites of Nyaa. Currently supported:
// - ProviderNyaa: http://nyaa.si
// - ProviderSukebei: http://sukebei.nyaa.si (NSFW)
type Provider int

const (
	ProviderNyaa Provider = iota
	ProviderSukebei
	providerEnd
)

func (p Provider) String() string {
	switch p {
	case ProviderNyaa:
		return "ProviderNyaa"
	case ProviderSukebei:
		return "ProviderSukebei"
	}

	return ""
}

func (p Provider) BaseURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   p.Host(),
	}
}

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
