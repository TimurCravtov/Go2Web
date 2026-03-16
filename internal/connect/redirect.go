package connect

import (
	"fmt"
	"slices"
)

// WithRedirects wraps a GetFunc to automatically follow HTTP redirects up to maxRedirects times.
func WithRedirects(next GetFunc, maxRedirects int) GetFunc {
	return func(url string, body []byte, headers map[string]string) (*HttpResponse, error) {
		currentURL := url

		for i := 0; i <= maxRedirects; i++ {
			resp, err := next(currentURL, body, headers)
			if err != nil {
				return nil, err
			}

			if slices.Contains([]int{301, 302, 303, 307, 308}, resp.StatusCode) {
				location, ok := resp.Headers["location"]
				if !ok {
					return resp, nil
				}
				
				currentURL = location
				continue
			}

			return resp, nil
		}

		return nil, fmt.Errorf("stopped after %d redirects", maxRedirects)
	}
}