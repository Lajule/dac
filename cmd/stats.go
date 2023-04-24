package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stats called")
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	// statsCmd.PersistentFlags().String("foo", "", "A help for foo")
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
