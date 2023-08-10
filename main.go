package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Lajule/dac/cmd"
	dac "github.com/Lajule/dac/context"
	"github.com/Lajule/dac/ent"
	_ "github.com/mattn/go-sqlite3"
)

const (
	EnvVarName = "DAC_DB_FILE"

	DefaultDBFileName = ".dac.db"
)

var (
	version = "dev"

	commit = "none"

	date = "unknown"
)

func main() {
	var dbFile string
	if val, ok := os.LookupEnv(EnvVarName); ok {
		dbFile = val
	} else {
		dbFile = DefaultDBFileName
	}

	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", dbFile))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := cmd.Execute(context.WithValue(context.Background(), dac.KeyName, dac.Value{
		Version: fmt.Sprintf("%s %s %s", version, commit, date),
		Client:  client,
	})); err != nil {
		log.Fatalf("failed executing command: %v", err)
	}
}
