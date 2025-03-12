package cmd

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	_ "modernc.org/sqlite"

	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/embedded"
	"github.com/spf13/cobra"
)

// TODO: only for testing purposes, delete me afterwards
var testDbCmd = &cobra.Command{
	Use:   "test-db",
	Short: "Sample cmd for test-driving sqlc",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		database, err := sql.Open("sqlite", ":memory:")
		if err != nil {
			return err
		}

		// create tables
		if _, err := database.ExecContext(ctx, embedded.DDL); err != nil {
			return err
		}

		queries := db.New(database)

		// list all authors
		authors, err := queries.ListAuthors(ctx)
		if err != nil {
			return err
		}
		log.Println(authors)

		// create author
		insertedAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
			Name: "anon",
			Bio:  sql.NullString{String: "pwnd you already"},
		})
		if err != nil {
			return err
		}
		log.Println(insertedAuthor)

		// get inserted author
		fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
		if err != nil {
			return err
		}

		log.Println("authors match?", reflect.DeepEqual(insertedAuthor, fetchedAuthor))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testDbCmd)
}
