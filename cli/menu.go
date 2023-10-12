package cli

import (
	"fmt"
	"project/entity"
	"project/handler"
)

func Run() {
	for {
		fmt.Println("=====================================")
		fmt.Println("Welcome to Sports Market")
		fmt.Println("=====================================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("9. Exit")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose menu: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			Login()
		case 2:
			register()
		case 9:
			fmt.Println("Terima kasih, sampai jumpa kembali")
			return
		default:
			fmt.Println("Wrong input")
		}
	}
}

func Login()  {
	var username, password string

	fmt.Println("=======LOGIN=======")
	fmt.Print("Masukkan username:")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password:")
	fmt.Scanln(&password)

	user, isAuthenticated := handler.Login(username, password)
	if isAuthenticated {
		fmt.Println("Login berhasil! selamat datang,", user.Username)
	} else {
		fmt.Println("Login gagal!")
	}
}


func register()  {
	var username, password string
	fmt.Println("=======REGISTRASI=======")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	dataRegist := entity.Users{
		Username: username,
		Password: password,
	}

	err := handler.Register(dataRegist)
	if err != nil {
		fmt.Println("Gagal Registrasi", dataRegist)
	} else {
		fmt.Println("Berhasil registrasi")
	}
}