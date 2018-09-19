package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/runrioter/bcbc/bank"
)

type (
	cardBinCheckResponse struct {
		CardType  string         `json:"cardType"`
		Bank      string         `json:"bank"`
		Key       string         `json:"key"`
		Messages  []errorMessage `json:"messages"`
		Validated bool           `json:"validated"`
		Stat      string         `json:"stat"`
		CardNo    string
	}

	errorMessage struct {
		Name       string `json:"name"`
		ErrorCodes string `json:"errorCodes"`
	}

	jsonOutput struct {
		Bin    string `json:"bin"`
		Bank   string `json:"bank"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Length int    `json:"length"`
	}
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
