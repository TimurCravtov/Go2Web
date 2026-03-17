package negociation


import (
	"strings"
	"go2web/internal/html"
)

func AcceptLanguages(langs []string) html.Headers {
	var parts []string
	for _, lang := range langs {
		if lang != "" {
			parts = append(parts, lang)
		}
	}
	value := strings.Join(parts, ", ")

	return html.Headers{
		"Accept-Language": value,
	}
}

func AcceptCharsets(charsets []string) html.Headers {
	return html.Headers{
		"Accept-Charset": strings.Join(charsets, ", "),
	}
}

func AcceptContentTypes(types []string) html.Headers {
	return html.Headers{
		"Accept": strings.Join(types, ", "),
	}
}