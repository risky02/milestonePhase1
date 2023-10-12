package handler

import (
	"context"
	"fmt"
	"project/config"
	"project/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register(dataRegist entity.Users) error {
	db, _ := config.InitDB()
	ctx := context.Background()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dataRegist.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	
	_, err = db.ExecContext(ctx, "INSERT INTO users (Username, Password) VALUES (?, ?)", dataRegist.Username, hashedPassword)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Register")
	return nil
}

