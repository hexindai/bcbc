package cmd

import (
	"fmt"

	"github.com/hexindai/bcbc/bank"
	"github.com/spf13/cobra"
)

func init() {
	bcbc.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all bank cardbins",
	Long:  "\nList all bank cardbins",
	Run: func(cmd *cobra.Command, args []string) {
		for _, bin := range bank.CardBINs {
			fmt.Printf("Bin: %s, Bank: %s, Name: %s Type: %s, Length: %v\n",
				bin.Bin,
				bin.Bank,
				bank.BankNameMap[bin.Bank],
				bin.Type,
				bin.Length,
			)
		}
	},
}
