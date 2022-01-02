package nyaa

import (
	"errors"
)

var (
	ErrUnknownProvider  = errors.New("unknown Provider")
	ErrUnknownFilterBy  = errors.New("unknown FilterBy")
	ErrUnknownSortBy    = errors.New("unknown SortBy")
	ErrUnknownSortOrder = errors.New("unknown SortOrder")
)
