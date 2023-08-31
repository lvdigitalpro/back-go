package adapters

import (
	"database/sql"
	"github.com/lvdigitalpro/back/src/infra/database"
	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/pq"
	"os"
)

func PgAdapter(db *sql.DB) database.DBInfra {

	open, err := apmsql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return open
}
