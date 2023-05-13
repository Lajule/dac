package cmd

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Lajule/dac/ui"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "dac",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			var input io.Reader

			if len(args) > 0 {
				file, err := os.Open(args[0])
				if err != nil {
					log.Fatal(err)
				}

				input = file
			} else {
				input = os.Stdin
			}

			b, err := io.ReadAll(input)
			if err != nil {
				log.Fatalf("failed reading input: %v", err)
			}

			if len(b) == 0 {
				log.Println("Input is empty")
				os.Exit(1)
			}

			app, err := ui.NewApp(10*time.Second, string(b))
			if err != nil {
				log.Fatalf("failed creating ui: %v", err)
			}

			app.Start()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed executing root command: %v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dac.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".dac")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file: ", viper.ConfigFileUsed())
	}
}
