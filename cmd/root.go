package cmd

import (
	"fmt"
	"os"

	"github.com/frandier/gocambridge/internal/cli"
	"github.com/spf13/cobra"
)

var (
	cookie string
	url    string
)

var rootCmd = &cobra.Command{
	Use:     "cambridge-go",
	Short:   "A CLI for solving Cambridge One exercises",
	Long:    `A CLI for solving Cambridge One exercises`,
	Example: `cambridge-go -c s%3AFKu-**** -u https://www.cambridgeone.org/nlp/apigateway/org_****/product/****/7e1e****`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if cookie == "" {
			cmd.Help()
			os.Exit(1)
		}
		if url == "" {
			cmd.Help()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.Run(url, cookie); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&cookie, "cookie", "c", "", "cookie")
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "url")
	rootCmd.MarkFlagsRequiredTogether("cookie", "url")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
