package cmd

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Lajule/dac/app"
	"github.com/Lajule/dac/ent"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "dac",
		Short: "Typing training sessions",
		Long:  `Dac is typing training sessions program, it's help you to improve your typing skills.`,
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
				log.Println("input is empty")
				os.Exit(1)
			}

			duration, err := cmd.Flags().GetDuration("duration")
			if err != nil {
				log.Fatalf("failed getting duration: %v", err)
			}

			close, err := cmd.Flags().GetBool("close")
			if err != nil {
				log.Fatalf("failed getting close: %v", err)
			}

			if close && duration == 0 {
				log.Println("a duration must be defined")
				os.Exit(1)
			}

			d, err := app.NewDac(duration, close, string(b))
			if err != nil {
				log.Fatalf("failed creating ui: %v", err)
			}

			d.Start()
			log.Println("training session: ", d)

			client, err := ent.Open("sqlite3", "file:dac.db?cache=shared&_fk=1")
			if err != nil {
				log.Fatalf("failed opening connection to sqlite: %v", err)
			}
			defer client.Close()
			if err := client.Schema.Create(context.Background()); err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}
			t, err := client.Training.
				Create().
				Save(context.Background())
			if err != nil {
				log.Fatalf("failed creating user: %v", err)
			}

			log.Println("training was created: ", t)
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
	rootCmd.Flags().DurationP("duration", "d", 0, "Duration of the training session")
	rootCmd.Flags().BoolP("close", "c", false, "Close on session timeout")
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
		log.Println("using config file: ", viper.ConfigFileUsed())
	}
}
