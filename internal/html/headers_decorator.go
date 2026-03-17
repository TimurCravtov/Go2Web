package html

import "go2web/internal/connect"

type Headers map[string]string

func WithHeaders(h Headers) func(connect.GetFunc) connect.GetFunc {
	return func(next connect.GetFunc) connect.GetFunc {
		return func(url string, body []byte, headers map[string]string) (*connect.HttpResponse, error) {
			headers = mergeHeaders(headers, h)
			return next(url, body, headers)
		}
	}
}

func mergeHeaders(base map[string]string, extra Headers) map[string]string {
	if base == nil {
		base = make(map[string]string)
	}
	for k, v := range extra {
		base[k] = v
	}
	return base
}
