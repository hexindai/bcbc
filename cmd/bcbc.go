package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// BCBCVERSION bcbc version
const BCBCVERSION = "0.0.8"

var bcbc = &cobra.Command{
	Use:   "bcbc",
	Short: "bcbc is a command for searching China's bankcard info",
	Long:  "\nbcbc is a command for searching China's bankcard info",
}

// Execute bcbc command
func Execute() {
	if err := bcbc.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
