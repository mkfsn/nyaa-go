package nyaa

import (
	"errors"
)

type FilterBy int

const (
	FilterByNoFilter FilterBy = iota
	FilterByNoRemakes
	FilterByTrustedOnly
	filterByEnd
)

func (f FilterBy) String() string {
	switch f {
	case FilterByNoFilter:
		return "NoFilter"
	case FilterByNoRemakes:
		return "NoRemakes"
	case FilterByTrustedOnly:
		return "TrustedOnly"
	}

	return ""
}

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
	if f >= filterByEnd {
		return errors.New("invalid FilterBy value")
	}

	return nil
}
