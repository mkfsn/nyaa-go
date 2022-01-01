package nyaa

import (
	"errors"
)

type SortBy int

const (
	SortByDate SortBy = iota
	SortByComments
	SortByDownloads
	SortBySeeders
	SortByLeechers
	SortBySize
	sortByEnd
)

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

	return ""
}

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
		return errors.New("invalid SortBy value")
	}

	return nil
}

type SortOrder int

const (
	SortOrderDesc SortOrder = iota
	SortOrderAsc
	sortOrderEnd
)

func (s SortOrder) String() string {
	switch s {
	case SortOrderDesc:
		return "Desc"
	case SortOrderAsc:
		return "Asc"
	}

	return ""
}

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
		return errors.New("invalid SortOrder value")
	}

	return nil
}
