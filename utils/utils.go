package utils

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectDatabase(url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// func AlwaysUseJsonMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Content-Type", "application/text; charset=utf-8")
// 	}
// }
