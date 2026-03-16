package html

import (
	"go2web/internal/connect"
	"mime"
	"net/http"
	"strings"
)

type ContentType string

const (
	TypeJSON       ContentType = "application/json"
	TypeHTML       ContentType = "text/html"
	TypePlainText  ContentType = "text/plain"
	TypeXML        ContentType = "application/xml"
	TypeForm       ContentType = "application/x-www-form-urlencoded"
	TypeMultipart  ContentType = "multipart/form-data"
	TypeJavascript ContentType = "application/javascript"
	TypeCSS        ContentType = "text/css"
)

func GetContentType(response *connect.HttpResponse) (ContentType, error) {
	// Prefer the Content-Type header when available.
	typeHeader, ok := response.Headers["content-type"]
	if ok && strings.TrimSpace(typeHeader) != "" {
		mediaType, _, err := mime.ParseMediaType(typeHeader)
		if err != nil {
			return "", err
		}
		return ContentType(mediaType), nil
	}

	// Fallback: sniff from body.
	detected := http.DetectContentType(response.Body)
	return ContentType(detected), nil
}
