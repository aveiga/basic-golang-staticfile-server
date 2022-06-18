package customdb

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customerrors"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customlogger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func GetDB() (*bun.DB, error) {
	logger := customlogger.NewCustomLogger()

	if os.Getenv("GIN_MODE") == models.PROD {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

		db := bun.NewDB(sqldb, pgdialect.New())

		return db, nil
	}

	if os.Getenv("GIN_MODE") == models.DEV {
		sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
		if err != nil {
			dbError := customerrors.RestError{
				Message: "Failed setting up DB access",
				Status:  http.StatusInternalServerError,
				Code:    "internal_server_error",
			}
			logger.Fatal(err)
			return nil, &dbError
		}

		db := bun.NewDB(sqldb, sqlitedialect.New())

		return db, nil
	}

	missingEnvError := customerrors.RestError{
		Message: "Missing GO ENV definition",
		Status:  http.StatusInternalServerError,
		Code:    "internal_server_error",
	}
	logger.Fatal(missingEnvError)
	return nil, &missingEnvError
}
