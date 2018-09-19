package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/hexindai/bcbc/bank"
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
	Long:  "\nSearch subcommand for searching bankcard info",
	Run: func(cmd *cobra.Command, args []string) {

		cbcr := bank.FetchCardBinCheckByCard(card)

		if output == "json" {
			cbcr.WriteResponse(os.Stdout, bank.JSONContentType)
		} else {
			cbcr.WriteResponse(os.Stdout, bank.TextContentType)
		}
	},
}
