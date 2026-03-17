package printer

import "go2web/internal/connect"

func TextPrinter(urlPath string, response *connect.HttpResponse) (string, error) {
	return string(response.Body), nil
}
