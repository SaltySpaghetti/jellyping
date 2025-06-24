package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDatabase() (*pgx.Conn, error) {
	url := GetDatabaseUrl()
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetDatabaseUrl() string {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresUser != "" && postgresPassword != "" && postgresDB != "" {
		return fmt.Sprintf("postgres://%s:%s@jellyping-db:5432/%s",
			postgresUser,
			postgresPassword,
			postgresDB,
		)
	}

	defaultDbName := "jellyping"
	defaultDbUser := "postgres"
	defaultDbPassword := "password"
	defaultDbHost := "jellyping-db"
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s",
		defaultDbUser,
		defaultDbPassword,
		defaultDbHost,
		defaultDbName,
	)

	return databaseUrl
}

// func AlwaysUseJsonMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Content-Type", "application/text; charset=utf-8")
// 	}
// }
