package main

import (
	bd "API/BD"
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:skj09@localhost:5432/postgres")
	if err != nil {
		fmt.Println("connection", err)
		return
	}
	connect := &bd.Connection{
		Connect: conn,
	}
	bd.CreateDB(conn, ctx)
	router := mux.NewRouter()
	router.Path("/registr").HandlerFunc(connect.Registr).Methods("POST")
	router.Path("/login").Methods("POST").HandlerFunc(connect.Login)
	fmt.Println("server is ready")
	if err := http.ListenAndServe(":9091", router); err != nil {
		fmt.Println("server", err)
		return
	}
}
