package bd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// мне надо юзера расшифровать и
func (conn *Connection) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		fmt.Println("decode", err)
	}
	sqlQuery := `
		SELECT email,password FROM newTask 
		WHERE email = $1
	`
	var passWord string
	var email string
	err = conn.Connect.QueryRow(r.Context(), sqlQuery, &userLogin.Email).Scan(&email, &passWord)
	if err != nil {
		fmt.Println("not email", err)
		json.NewEncoder(w).Encode("not email")
	} else {
		json.NewEncoder(w).Encode("email is ready")
	}
	if userLogin.Password == passWord {
		json.NewEncoder(w).Encode("вошел")
	} else {
		json.NewEncoder(w).Encode("неверный пароль")
	}

}
