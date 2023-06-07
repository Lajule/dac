package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Lajule/dac/cmd"
	"github.com/Lajule/dac/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dbFile string
	if val, ok := os.LookupEnv("DAC_DB_FILE"); ok {
		dbFile = val
	} else {
		dbFile = "dac.db"
	}

	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", dbFile))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := cmd.Execute(context.WithValue(context.Background(), "client", client)); err != nil {
		log.Fatalf("failed executing command: %v", err)
	}
}
