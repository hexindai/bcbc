package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hexindai/bcbc/bank"
	"github.com/spf13/cobra"
)

func init() {
	bcbc.AddCommand(randomCmd)
	rand.Seed(time.Now().UnixNano())
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Return a random bankcard",
	Long:  "\nReturn a random bankcard",
	Run: func(cmd *cobra.Command, args []string) {
		bin := randomBin()
		bankcard := make([]byte, 0, bin.Length)
		bankcard = append(bankcard, bin.Bin...)
		remaining := bin.Length - len(bin.Bin)
		numbers := "0123456789"
		for i := 0; i < remaining; i++ {
			bankcard = append(bankcard, byte(numbers[rand.Intn(10)]))
		}
		fmt.Printf(
			"BIN: %s, Type: %s, Length: %v, Bank: %s, BankName: %s, No: %s\n",
			bin.Bin,
			bin.Type,
			bin.Length,
			bin.Bank,
			bin.BankName(),
			string(bankcard),
		)
	},
}

func randomBin() *bank.CardBIN {
	l := len(bank.CardBINs)
	i := rand.Intn(l)
	return bank.CardBINs[i]
}
