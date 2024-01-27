package main

import (
	"context"
	"flag"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

func main() {
	flags := flag.NewFlagSet("goose", flag.ExitOnError)
	dir := flags.String("dir", "./migrations", "directory with migration files")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	args := flags.Args()
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]
	dbstring := "file:db/local.sqlite?_foreign_keys=on"
	db, err := goose.OpenDBWithDriver("sqlite3", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}
	_, err = db.Exec("PRAGMA foriegn_keys = ON")
	if err != nil {
		log.Fatalf("goose: failed to set foreign_key constraint %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()
	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}
	if command == "create" {
		arguments = append(arguments, "sql")
	}

	if err := goose.RunContext(context.Background(), command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
