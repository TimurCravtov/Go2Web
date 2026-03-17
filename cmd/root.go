package cmd

import (
	"fmt"
	"os"
	"go2web/internal/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go2web",
	Short: "Go2Web is a CLI tool for searching the web and retrieving results.",
	Long: `Go2Web is a command-line application that allows users to search the web using various search engines and retrieve results directly in the terminal. It supports features like caching, redirects, and customizable search engines.`,
	Run: func(cmd *cobra.Command, args []string) {

		if searchQuery, _ := cmd.Flags().GetString("search"); searchQuery != "" {
			if dynamic, _ := cmd.Flags().GetBool("dynamic"); dynamic {
				cli.HandleSearchDynamic(cmd, args)
			} else {
				cli.HandleSearch(cmd, args)
			}
		}
		if url, _ := cmd.Flags().GetString("url"); url != "" {
			cli.HandleUrlRequest(cmd, args)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringP("search", "s", "", "Search the web with a query")
	rootCmd.Flags().BoolP("dynamic", "d", false, "Use dynamic search (with CLI interface)")

	rootCmd.Flags().StringP("engine", "e", "startpage", "Search engine to use (e.g., startpage, mojeek)")
	rootCmd.Flags().StringP("url", "u", "", "Fetch and display the content of a URL")
	rootCmd.Flags().IntP("max-redirects", "", 10, "Maximum number of redirects to follow when fetching a URL. Pass -1 to not limit redirects. Default is 10.")


	// content negotiation
	rootCmd.Flags().StringArrayP("lang", "l", []string{"en"}, "List of accepted languages for content negotiation (e.g., en, fr, es)")
	rootCmd.Flags().StringArrayP("charset", "c", []string{"UTF-8"}, "List of accepted charsets for content negotiation (e.g., UTF-8, ISO-8859-1)")
	rootCmd.Flags().StringArrayP("type", "t", []string{"*/*"}, "List of accepted content types for content negotiation (e.g., application/json, text/plain)")

	// and mark them only for url
	OnlyValidWith(rootCmd, "lang", "url")
	OnlyValidWith(rootCmd, "charset", "url")
	OnlyValidWith(rootCmd, "type", "url")

	rootCmd.MarkFlagsMutuallyExclusive("search", "url") // you can't use search and url together
	rootCmd.Flags().BoolP("no-cache", "", false, "Disable caching")

	OnlyValidWith(rootCmd, "engine", "search") // you can only use engine if search is provided
	OnlyValidWith(rootCmd, "max-redirects", "url") // you can only use max-redirects if url is provided

}

func OnlyValidWith(cmd *cobra.Command, dependentFlag, requiredFlag string) error {
	dependentProvided := cmd.Flags().Changed(dependentFlag)
	requiredProvided := cmd.Flags().Changed(requiredFlag)

	if dependentProvided && !requiredProvided {
		return fmt.Errorf("error: the --%s flag is only valid when --%s is also provided", dependentFlag, requiredFlag)
	}
	return nil
}

