package bookData

import (
	"context"
	"fmt"
	"log"
	"scylladb/constants"
	"scylladb/database"
	"scylladb/graph/model"

	"github.com/gocql/gocql"
)

const (
	keyspaceName = "bookData"
	tableName    = "book"
)

func CreateBook(ctx context.Context, input model.CreateBookListingInput) (*model.BookListing, error) {
	session, err := database.Connect()
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorMessage["DatabaseConnectionFailed"], "%s", err)
	}

	log.Println(constants.StringCOnstants["DatabaseConnected"])
	// Create a keyspace if it doesn't exist
	keyspaceQuery := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};", keyspaceName)
	if err := session.Query(keyspaceQuery).Exec(); err != nil {
		return nil, fmt.Errorf(constants.ErrorMessage["KeyspaceFailed"], "%s", err)
	}
	// Create a table if it doesn't exist
	tableQuery := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s.%s (
            id UUID PRIMARY KEY,
            title TEXT,
            author TEXT,
            description TEXT
        );
    `, keyspaceName, tableName)
	if err := session.Query(tableQuery).Exec(); err != nil {
		return nil, fmt.Errorf(constants.ErrorMessage["TableFailed"], "%s", err)
	}
	// Generate a new ID for the book listing
	bookID := gocql.TimeUUID()
	// Create a new book listing object with the input data
	createdBook := &model.BookListing{
		ID:          bookID.String(),
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
	}
	// Insert the book listing into the "book" table in the database
	// query := session.Query(`
	//     INSERT INTO %s.%s (id, title, author, description)
	//     VALUES (?, ?, ?, ?)
	// `, keyspaceName, tableName)

	insertQuery := fmt.Sprintf(`
        INSERT INTO %s.%s (id, author, title, description)
        VALUES (?, ?, ?, ?);
    `, keyspaceName, tableName)

	query := session.Query(insertQuery, createdBook.ID, createdBook.Author, createdBook.Title, createdBook.Description)

	if err := query.Exec(); err != nil {
		return nil, fmt.Errorf(constants.ErrorMessage["FailedToInsetData"], "%s", err)
	}

	return createdBook, nil
}
