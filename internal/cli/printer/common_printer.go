package printer

import (
    "fmt"
    "net/url"
    "regexp"
	"github.com/0magnet/calvin"
	"go2web/internal/cli/printer/utils"
    _ "image/jpeg"
    _ "image/png"
	"go2web/internal/request"
    "strings"

)

func WithHero(next HttpResponsePrinter) HttpResponsePrinter {
    return func(url string, response *request.HttpResponse) (string, error) {
        hero := buildWebsiteHero(response, url)
        content, err := next(url, response)
        return hero + content, err
    }
}

func WithHeaders(next HttpResponsePrinter) HttpResponsePrinter {
    return func(url string, response *request.HttpResponse) (string, error) {
        var sb strings.Builder
        headers := response.Headers
        for key, value := range headers {
            sb.WriteString(fmt.Sprintf("%s: %s\n", utils.Colorize(key, utils.ColorMagenta), value))
        }
        
        nextResponse, err := next(url, response)
        if err != nil {
            return "", err
        }
        return sb.String() + "\n" + nextResponse, nil
    }
}

func WithStatusLine(next HttpResponsePrinter) HttpResponsePrinter {
    return func(url string, response *request.HttpResponse) (string, error) {
        // Determine color based on the first digit of the status code
        var statusColor string
        switch response.StatusCode / 100 {
        case 2:
            statusColor = utils.ColorGreen   // Success
        case 3:
            statusColor = utils.ColorBlue    // Redirects
        case 4:
            statusColor = utils.ColorYellow  // Client Errors
        case 5:
            statusColor = utils.ColorRed     // Server Errors
        default:
            statusColor = utils.ColorReset   // Fallback for 1xx or unknown
        }

        rawStatus := fmt.Sprintf("%d %s", response.StatusCode, response.StatusText)
        coloredStatus := utils.Colorize(rawStatus, statusColor)
        
        statusLine := fmt.Sprintf("Status: %s", coloredStatus)

        nextResponse, err := next(url, response)
        if err != nil {
            return "", err
        }
        
        return statusLine + "\n" + nextResponse, nil
    }
}

func buildWebsiteHero(response *request.HttpResponse, rootUrl string) string {
    faviconUrl := getFavicoLink(response, rootUrl)

	resp, _ := request.Get(faviconUrl, nil, nil)
	
    asciiFavicon, err := utils.ImageToAscii(resp.Body, 12, 12)
	if err != nil {
		asciiFavicon = ""
	}
    var sb strings.Builder

    u, _ := url.Parse(rootUrl)

	websiteName := u.Hostname() 
    
    asciiTitle := calvin.AsciiFont(strings.ToUpper(websiteName))
    titleLines := strings.Split(strings.TrimRight(asciiTitle, "\n"), "\n")

    var iconLines []string
    boxWidth := 24 
    
    if err == nil {
        rawIconLines := strings.Split(strings.TrimRight(asciiFavicon, "\n"), "\n")
        
        iconLines = append(iconLines, "╭"+strings.Repeat("─", boxWidth + 2)+"╮")
        for _, line := range rawIconLines {
            iconLines = append(iconLines, "│ "+line+" │")
        }
        iconLines = append(iconLines, "╰"+strings.Repeat("─", boxWidth + 2)+"╯")
    }

    iconHeight := len(iconLines)
    titleHeight := len(titleLines)
    
    maxLines := iconHeight
    if titleHeight > maxLines {
        maxLines = titleHeight
    }

    // Calculate vertical starting offsets for true centering
    iconOffset := (maxLines - iconHeight) / 2
    titleOffset := (maxLines - titleHeight) / 2

    emptyIconPadding := strings.Repeat(" ", boxWidth)

    for i := 0; i < maxLines; i++ {
        // Determine the icon row (or pad with spaces if above/below the icon bounds)
        iconPart := emptyIconPadding
        if len(iconLines) > 0 && i >= iconOffset && i < iconOffset+iconHeight {
            iconPart = iconLines[i-iconOffset]
        } else if len(iconLines) == 0 {
            iconPart = ""
        }

        titlePart := ""
        if i >= titleOffset && i < titleOffset+titleHeight {
            titlePart = titleLines[i-titleOffset]
        }

        if iconPart != "" && titlePart != "" {
            sb.WriteString(iconPart + "   " + titlePart + "\n")
        } else if iconPart != "" {
            sb.WriteString(iconPart + "\n")
        } else {
            sb.WriteString(titlePart + "\n")
        }
    }

    return sb.String()
}

func getFavicoLink(response *request.HttpResponse, rootUrl string) string {
    baseURL, err := url.Parse(rootUrl)
    if err != nil {
        return ""
    }

    bodyStr := string(response.Body)

    linkRegex := regexp.MustCompile(`(?i)<link[^>]+>`)
    hrefRegex := regexp.MustCompile(`(?i)href\s*=\s*["']([^"']+)["']`)
    relRegex := regexp.MustCompile(`(?i)rel\s*=\s*["']([^"']+)["']`)

    links := linkRegex.FindAllString(bodyStr, -1)

    for _, linkTag := range links {
        relMatch := relRegex.FindStringSubmatch(linkTag)
        if len(relMatch) > 1 {
            relVal := strings.ToLower(relMatch[1])

            if strings.Contains(relVal, "icon") {
                hrefMatch := hrefRegex.FindStringSubmatch(linkTag)
                if len(hrefMatch) > 1 {
                    rawHref := hrefMatch[1]

                    hrefURL, err := url.Parse(rawHref)
                    if err == nil {
                        // Resolve relative URLs against the base URL
                        resolvedURL := baseURL.ResolveReference(hrefURL)
                        return resolvedURL.String()
                    }
                }
            }
        }
    }

    fallbackURL, _ := url.Parse("/favicon.ico")
    resolvedFallback := baseURL.ResolveReference(fallbackURL)

    return resolvedFallback.String()
}
