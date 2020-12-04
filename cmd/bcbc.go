package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// BCBCVERSION bcbc version
const BCBCVERSION = "0.0.10"

var bcbc = &cobra.Command{
	Use:   "bcbc",
	Short: "China UnionPay Bank Card BIN Checker",
	Long:  "China UnionPay Bank Card BIN Checker",
}

// Execute bcbc command
func Execute() {
	if err := bcbc.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
