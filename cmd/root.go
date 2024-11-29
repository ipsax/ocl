package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ocl",
	Short: "ocl is a tool for learning OpenCypher",
	Long:  `ocl (OpenCypher Learning) is a tool for learning OpenCypher with the PostgreSQL Apache Age extension"`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
