package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hexindai/bcbc/bank"

	"github.com/spf13/cobra"
)

func init() {
	bcbc.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all bank card BINs",
	Long:  "\nList all bank card BINs",
	Run: func(cmd *cobra.Command, args []string) {

		pager := os.Getenv("PAGER")

		if pager == "" {
			pager = "less"
		}

		c := exec.Command(pager)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Env = os.Environ()

		stdin, err := c.StdinPipe()
		if err != nil {
			log.Fatalln(err)
		}
		go func() {
			defer stdin.Close()
			for _, bin := range bank.CardBINs {
				fmt.Fprintf(stdin, "Bin: %s, Bank: %s, Name: %s Type: %s, Length: %v\n",
					bin.Bin,
					bin.Bank,
					bin.BankName(),
					bin.Type,
					bin.Length,
				)
			}
		}()

		err = c.Run()
		if err != nil {
			log.Fatalln(err)
		}

	},
}
