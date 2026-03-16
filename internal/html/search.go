package html

import "go2web/internal/connect"

type SearchResult struct {
	Title   string
	URL     string
	Snippet string
}

type Search interface {
	Search(query string, page int, get connect.GetFunc) ([]SearchResult, error)
}
