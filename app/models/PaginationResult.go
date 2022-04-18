package models

type PaginationResult struct {
	Data        interface{}
	Total       int
	PerPage     int
	CurrentPage int
}
