package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	dac "github.com/Lajule/dac/context"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Program version",
		Long:  `Display program version.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			val := ctx.Value(dac.KeyName).(dac.Value)
			fmt.Println(val.Version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
