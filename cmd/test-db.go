package cmd

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	_ "modernc.org/sqlite"

	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/spf13/cobra"
)

// TODO: only for testing purposes, delete me afterwards
func NewTestDbCmd(connection db.DbConnection) CobraCommand {
	testDbCmd := &cobra.Command{
		Use:   "test-db",
		Short: "Sample cmd for test-driving sqlc",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			q := queries.New(connection.Database())

			// list all authors
			authors, err := q.ListAuthors(ctx)
			if err != nil {
				return err
			}
			log.Println(authors)

			// create author
			insertedAuthor, err := q.CreateAuthor(ctx, queries.CreateAuthorParams{
				Name: "anon",
				Bio:  sql.NullString{String: "pwnd you already"},
			})
			if err != nil {
				return err
			}
			log.Println(insertedAuthor)

			// get inserted author
			fetchedAuthor, err := q.GetAuthor(ctx, insertedAuthor.ID)
			if err != nil {
				return err
			}

			log.Println("authors match?", reflect.DeepEqual(insertedAuthor, fetchedAuthor))

			return nil
		},
	}

	return GenericCommand{
		cmd:  testDbCmd,
		path: "root test-db",
	}
}
