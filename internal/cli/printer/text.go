package printer

import "go2web/internal/request"

func TextPrinter(urlPath string, response *request.HttpResponse) (string, error) {
	return string(response.Body), nil
}
