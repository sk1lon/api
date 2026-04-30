package bd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (conn *Connection) Registr(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("decode", err)
		return
	}
	sqlQuery := `
		INSERT INTO newTask (email, password)
		VALUES ($1,$2)
	`
	_, err = conn.Connect.Exec(r.Context(), sqlQuery, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("to db", err)
		json.NewEncoder(w).Encode("ошибка")
	} else {
		json.NewEncoder(w).Encode("данные сохранены")
	}

}
