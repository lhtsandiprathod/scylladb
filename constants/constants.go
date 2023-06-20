package constants

var (
	ErrorMessage = map[string]string{
		"GetBookError":                 "Error while fetching books",
		"CreateBookError":              "Error while creating books",
		"UpdateBookByIdError":          "Error while updating books by ID",
		"UpdateBooksByAuthorNameError": "Error while updating books by author name",
		"DeleteBookError":              "Error while deleting book",
		"KeyspaceFailed":               "Failed to create keyspace",
		"TableFailed":                  "Failed to create table",
		"FailedToInsetData":            "Failed to insert data",
	}

	SuccessMessage = map[string]string{
		"BookFetched":             "Book fetched successfully",
		"BooksFetched":            "Books fetched successfully",
		"BookUpdatedById":         "Book successfully updated by authorName",
		"BookUpdatedByAuthorName": "Book successfully updated by Id",
		"BookDeleted":             "Book deleted successfully",
		"BookCreated":             "Book created successfully",
	}

	StringCOnstants = map[string]string{
		"ServerURL":                "mongodb://localhost:27017",
		"DatabaseConnected":        "Database connected successfully",
		"DatabaseConnectionFailed": "failed to connect to the database:",
	}
)
