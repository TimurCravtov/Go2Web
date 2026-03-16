package cli

import (
	"fmt"
	"go2web/internal/connect"
	"go2web/internal/html"
	"go2web/internal/html/search_engines"
	"strings"

	"github.com/0magnet/calvin"
	"github.com/spf13/cobra"
)

func HandleSearch(cmd *cobra.Command, args []string) {
	searchQuery, _ := cmd.Flags().GetString("search")
	if searchQuery == "" {
		return
	}

	// select engine
	engineName, _ := cmd.Flags().GetString("engine")
	var engine html.Search

	switch engineName {
	case "startpage":
		engine = search_engines.NewStartpageSearchEngine("https://www.startpage.com/sp/search?query=")
	case "mojeek":
		engine = search_engines.NewMojeekSearchEngine("https://www.mojeek.com/search?q=")
	default:
		fmt.Printf("Unknown search engine: %s\n", engineName)
		return
	}

	// cute input logo
	fmt.Println(calvin.AsciiFont(strings.ToUpper(engineName)))
	fmt.Println("┌───────────────────────────────────────────────┐")
	fmt.Printf("│ %-43s ⌕ │\n", searchQuery)
	fmt.Println("└───────────────────────────────────────────────┘")

	// execute
	results, err := engine.Search(searchQuery, 1, connect.Get)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// print results
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result.Title)
		fmt.Printf("   URL: %s\n", html.Colorize(result.URL, html.ColorBlue))
		fmt.Println("   " + "─" + "─" + "─" + "─" + "─" + "─" + "─" + "─" + "─" + "─" + "─" + "─" + "─")
	}
}