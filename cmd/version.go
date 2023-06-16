package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "development"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Program version",
		Long:  `Display program version.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
