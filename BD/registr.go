package bd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err = conn.Connect.Exec(r.Context(), sqlQuery, user.Email, hashPass)
	if err != nil {
		fmt.Println("to db", err)
		http.Error(w, "ошибка", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("данные сохранены")

}
