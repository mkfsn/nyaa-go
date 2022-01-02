package nyaa

// SortBy represents the way to sort the torrent when searching.
type SortBy int

// The ways to sort the torrents.
const (
	SortByDate SortBy = iota
	SortByComments
	SortByDownloads
	SortBySeeders
	SortByLeechers
	SortBySize
	sortByEnd
)

// String implements fmt.Stringer interface.
func (s SortBy) String() string {
	switch s {
	case SortByDate:
		return "Date"
	case SortByComments:
		return "Comments"
	case SortByDownloads:
		return "Downloads"
	case SortBySeeders:
		return "Seeders"
	case SortByLeechers:
		return "Leechers"
	case SortBySize:
		return "Size"
	}

	return unknownEntityName
}

// Value returns the value of the query parameter in the HTTP request.
func (s SortBy) Value() string {
	switch s {
	case SortByDate:
		return "id"
	case SortByComments:
		return "comments"
	case SortByDownloads:
		return "downloads"
	case SortBySeeders:
		return "seeders"
	case SortByLeechers:
		return "leechers"
	case SortBySize:
		return "size"
	}

	return ""
}

func (s SortBy) validate() error {
	if s >= sortByEnd {
		return ErrUnknownSortBy
	}

	return nil
}

// SortOrder represents the order of the torrent when searching.
type SortOrder int

// The order of the sorting.
const (
	SortOrderDesc SortOrder = iota
	SortOrderAsc
	sortOrderEnd
)

// String implements fmt.Stringer interface.
func (s SortOrder) String() string {
	switch s {
	case SortOrderDesc:
		return "Desc"
	case SortOrderAsc:
		return "Asc"
	}

	return unknownEntityName
}

// Value returns the value of the query parameter in the HTTP request.
func (s SortOrder) Value() string {
	switch s {
	case SortOrderDesc:
		return "desc"
	case SortOrderAsc:
		return "asc"
	}

	return ""
}

func (s SortOrder) validate() error {
	if s >= sortOrderEnd {
		return ErrUnknownSortOrder
	}

	return nil
}
