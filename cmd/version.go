package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	bcbc.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version and exit",
	Long:  "\nPrint version and exit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("bcbc version %s\n", BCBCVERSION)
	},
}
