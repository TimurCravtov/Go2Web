package search_engines

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"go2web/internal/request"
	"go2web/internal/html"

	"github.com/PuerkitoBio/goquery"
)

type DuckDuckGoSearchEngine struct {
	searchURL string
}

func NewDuckDuckGoSearchEngine(searchURL string) *DuckDuckGoSearchEngine {
	return &DuckDuckGoSearchEngine{searchURL: searchURL}
}

func (d *DuckDuckGoSearchEngine) Search(query string, page int, get request.GetFunc) ([]html.SearchResult, error) {
	var headers = map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language": "en-US,en;q=0.5",
	}

	// Assuming searchURL is initialized as "https://html.duckduckgo.com/html/?q="
	reqUrl := d.searchURL + url.QueryEscape(query)
	
	// DuckDuckGo uses an offset ('s') or page parameter ('p') depending on the endpoint.
	// For the HTML endpoint, usually appending p handles pagination.
	if page > 1 {
		reqUrl += fmt.Sprintf("&p=%d", page)
	}

	res, err := get(reqUrl, nil, headers)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body))
	if err != nil {
		return nil, err
	}

	var results []html.SearchResult
	
	// DuckDuckGo's HTML version groups results in the '.result' class
	doc.Find(".result").Each(func(i int, sel *goquery.Selection) {
		titleSel := sel.Find("h2.result__title a.result__a")
		title := strings.TrimSpace(titleSel.Text())
		
		href, _ := titleSel.Attr("href")
		href = strings.TrimSpace(href)

		// DuckDuckGo wraps outbound links in a redirect tracker. 
		// We extract the 'uddg' query parameter to get the clean target URL.
		if strings.Contains(href, "uddg=") {
			u, err := url.Parse(href)
			if err == nil {
				cleanURL := u.Query().Get("uddg")
				if cleanURL != "" {
					href = cleanURL
				}
			}
		} else if strings.HasPrefix(href, "//") {
			href = "https:" + href
		}

		snippet := strings.TrimSpace(sel.Find("a.result__snippet").Text())

		if title != "" && href != "" && strings.HasPrefix(href, "http") {
			results = append(results, html.SearchResult{
				Title:   title,
				URL:     href,
				Snippet: snippet,
			})
		}
	})

	return results, nil
}