package cli

import (
	"fmt"
	"go2web/internal/connect"
	"go2web/internal/html"
	"go2web/internal/printer"
	"math"
	"strings"

	_ "github.com/mat/besticon/ico"
	"github.com/spf13/cobra"
)

func HandleUrlRequest(cmd *cobra.Command, args []string) {

    rawURL, _ := cmd.Flags().GetString("url")
    urlStr := rawURL
    if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
        urlStr = "https://" + urlStr
    }

    var getter connect.GetFunc = connect.Get
    noCache, _ := cmd.Flags().GetBool("no-cache")
    if !noCache {
        cache := connect.NewFileCache("cache")
        getter = cache.WithCache(getter)
    }

    redirectCount, _ := cmd.Flags().GetInt("max-redirects")
    if redirectCount < 0 {
        redirectCount = math.MaxInt
    }
    if redirectCount >= 0 {
        getter = connect.WithRedirects(getter, redirectCount)
    }

    response, err := getter(urlStr, nil, nil)
    if err != nil {
        fmt.Printf("Error fetching page: %v\n", err)
        return
    }

    var basePrinter printer.HttpResponsePrinter

    contentType, err := html.GetContentType(response)

    if err != nil {
        fmt.Printf("Error determining content type: %v\n", err)
        return
    }

    switch contentType {
    case html.TypeHTML:
        basePrinter = printer.HtmlResponseParser
    case html.TypeJSON:
        basePrinter = printer.JsonPrinter
    default:
        basePrinter = printer.HtmlResponseParser
    }

    printer := printer.WithHeaders(printer.WithHero(basePrinter))
    
    str, _ := printer(urlStr, response);
    
    fmt.Println(str)

}


