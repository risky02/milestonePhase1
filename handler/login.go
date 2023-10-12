package handler

import (
	"project/config"
	"project/entity"

	"golang.org/x/crypto/bcrypt"
)


func Login(username, password string) (entity.Users, bool){
	db, _ := config.InitDB()
	var user entity.Users
	
	row := db.QueryRow("SELECT Username, Password FROM users WHERE Username=?", username)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		return user, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}

	return user, true
}