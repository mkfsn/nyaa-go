package nyaa

// FilterBy represents a way to filter the torrents when searching.
type FilterBy int

// The current ways for filtering torrents.
const (
	FilterByNoFilter FilterBy = iota
	FilterByNoRemakes
	FilterByTrustedOnly
	filterByEnd
)

// String implements fmt.Stringer interface.
func (f FilterBy) String() string {
	switch f {
	case FilterByNoFilter:
		return "NoFilter"
	case FilterByNoRemakes:
		return "NoRemakes"
	case FilterByTrustedOnly:
		return "TrustedOnly"
	}

	return unknownEntityName
}

// Value returns the value of the query parameter in the HTTP request.
func (f FilterBy) Value() string {
	switch f {
	case FilterByNoFilter:
		return "0"
	case FilterByNoRemakes:
		return "1"
	case FilterByTrustedOnly:
		return "2"
	}

	return ""
}

func (f FilterBy) validate() error {
	if f < 0 || f >= filterByEnd {
		return ErrUnknownFilterBy
	}

	return nil
}
