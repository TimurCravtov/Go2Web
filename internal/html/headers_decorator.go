package html

import (
    "go2web/internal/request"
    "net/http"
)

type Headers map[string]string

func WithHeaders(h Headers) func(request.GetFunc) request.GetFunc {
	return func(next request.GetFunc) request.GetFunc {
		return func(url string, body []byte, headers map[string]string) (*request.HttpResponse, error) {
			headers = mergeHeaders(headers, h)
			return next(url, body, headers)
		}
	}
}

func mergeHeaders(base map[string]string, extra Headers) map[string]string {
    merged := make(map[string]string)
    
    for k, v := range base {
        merged[k] = v
    }

    for k, v := range extra {
        canonicalKey := http.CanonicalHeaderKey(k)
        merged[canonicalKey] = v
    }  
    
    return merged
}