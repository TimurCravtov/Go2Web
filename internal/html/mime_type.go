package html

import (
	"go2web/internal/request"
	"mime"
	"net/http"
	"strings"
)

type ContentType string

const (
	TypeJSON      ContentType = "application/json"
	TypeHTML      ContentType = "text/html"
	TypePlainText ContentType = "text/plain"
	TypeXML       ContentType = "application/xml"
	TypeForm      ContentType = "application/x-www-form-urlencoded"
	TypeMultipart ContentType = "multipart/form-data"
	TypeJavascript ContentType = "application/javascript"
	TypeCSS       ContentType = "text/css"

	TypePNG       ContentType = "image/png"
	TypeJPEG      ContentType = "image/jpeg"
	TypeGIF       ContentType = "image/gif"
)

func (c ContentType) IsImage() bool {
	return strings.HasPrefix(string(c), "image/")
}

func GetContentType(response *request.HttpResponse) (ContentType, error) {
	typeHeader, ok := response.Headers["content-type"]
	if !ok {

		typeHeader, ok = response.Headers["content-type"]
	}

	if ok && strings.TrimSpace(typeHeader) != "" {
		mediaType, _, err := mime.ParseMediaType(typeHeader)
		if err != nil {
			return "", err
		}
		return ContentType(mediaType), nil
	}

	detected := http.DetectContentType(response.Body)
	mediaType, _, _ := mime.ParseMediaType(detected)
	return ContentType(mediaType), nil
}