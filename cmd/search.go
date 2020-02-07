package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hexindai/bcbc/response"
)

var (
	card   string
	output string
)

func init() {
	bcbc.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&card, "card", "c", "", "Bank `card number` to be checked")
	searchCmd.Flags().StringVarP(&output, "output", "o", "text", "Output format: `text, json`")
	searchCmd.MarkFlagRequired("card")
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search bankcard info",
	Long:  "Search bankcard info using card number",
	Run: func(cmd *cobra.Command, args []string) {

		cbcr := response.New(card)

		if output == "json" {
			cbcr.WriteResponse(os.Stdout, response.JSONContentType)
		} else {
			cbcr.WriteResponse(os.Stdout, response.TextContentType)
		}
		fmt.Println()

	},
}
