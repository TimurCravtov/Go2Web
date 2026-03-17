package printer

import "go2web/internal/request"

type HttpResponsePrinter func(url string, response *request.HttpResponse) (string, error)
