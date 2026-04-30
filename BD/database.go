package bd

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateDB(conn *pgx.Conn, ctx context.Context) {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS newTask (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE, 
		password TEXT UNIQUE
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		fmt.Println("db", err)
		return
	}
}
