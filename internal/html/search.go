package html

import "go2web/internal/request"

type SearchResult struct {
	Title   string
	URL     string
	Snippet string
}

type Search interface {
	Search(query string, page int, get request.GetFunc) ([]SearchResult, error)
}
