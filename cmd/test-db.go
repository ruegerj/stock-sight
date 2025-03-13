package cmd

import (
	"context"
	"log"
	"reflect"

	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/spf13/cobra"
)

// TODO: only for testing purposes, delete me afterwards
func NewTestDbCmd(ctx context.Context, authorRepo repository.AuthorRepository) CobraCommand {
	testDbCmd := &cobra.Command{
		Use:   "test-db",
		Short: "Sample cmd for test-driving sqlc",
		RunE: func(cmd *cobra.Command, args []string) error {
			// list all authors
			authors, err := authorRepo.GetAll(ctx)
			if err != nil {
				return err
			}
			log.Println(authors)

			// create author
			var bio string = "pwnd you already"
			insertedAuthor, err := authorRepo.Create(ctx, "anon", &bio)
			if err != nil {
				return err
			}
			log.Println(insertedAuthor)

			// get inserted author
			fetchedAuthor, err := authorRepo.GetById(ctx, insertedAuthor.ID)
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
