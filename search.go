package nyaa

import (
	"net/url"
)

// SearchOptions represents the options for searching torrents.
type SearchOptions struct {
	Provider  Provider
	FilterBy  FilterBy
	Category  Category
	Query     string
	SortBy    SortBy
	SortOrder SortOrder
}

func (o SearchOptions) validate() error {
	if err := o.Provider.validate(); err != nil {
		return err
	}

	if err := o.FilterBy.validate(); err != nil {
		return err
	}

	if err := o.Category.validate(o.Provider); err != nil {
		return err
	}

	if err := o.SortBy.validate(); err != nil {
		return err
	}

	return o.SortOrder.validate()
}

func (o SearchOptions) buildURL() *url.URL {
	return &url.URL{
		Scheme:   "https",
		Host:     o.Provider.Host(),
		RawQuery: o.buildQuery().Encode(),
	}
}

func (o SearchOptions) buildQuery() url.Values {
	values := make(url.Values)

	values.Set("f", o.FilterBy.Value())           // FilterBy
	values.Set("c", o.Category.Value(o.Provider)) // Category
	values.Set("q", o.Query)
	values.Set("s", o.SortBy.Value())
	values.Set("o", o.SortOrder.Value())

	return values
}
