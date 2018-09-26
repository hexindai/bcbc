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
	Short: "print version and exit",
	Long:  "\nprint bcbc version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("bcbc version %s\n", BCBCVERSION)
	},
}
