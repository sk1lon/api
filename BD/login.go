package bd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// мне надо юзера расшифровать и
func (conn *Connection) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		fmt.Println("decode", err)
	}
	sqlQuery := `
		SELECT password FROM newTask 
		WHERE email = $1
	`
	var passWord string
	err = conn.Connect.QueryRow(r.Context(), sqlQuery, userLogin.Email).Scan(&passWord)
	if err != nil {
		json.NewEncoder(w).Encode("NOT TRUE POCHTA")
		fmt.Println("NOT POCHTA")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(passWord), []byte(userLogin.Password))
	if err != nil {
		json.NewEncoder(w).Encode("NOT TRUE PASS")
		return
	}
	fmt.Println("ALL GOOF")
}
