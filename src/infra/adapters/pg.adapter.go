package adapters

import (
	"database/sql"
	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/pq"
	"os"
)

func PgAdapter() *sql.DB {

	open, err := apmsql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return open
}
