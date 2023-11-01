package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
)

var (
	PgDB       *sql.DB
	DBHost     = "database"
	DBPort     = "5432"
	DBUser     = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName     = os.Getenv("POSTGRES_DB")
)

func initializeDatabase() {
	connInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable", DBHost, DBUser, DBPassword, DBName,
	)
	var err error
	PgDB, err = sql.Open("pgx", connInfo)

	if err != nil {
		panic(err)
	}

	defer PgDB.Close()

	if err = PgDB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database is running on port 5432")
}
