package bd

import "github.com/jackc/pgx/v5"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Connection struct {
	Connect *pgx.Conn
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
