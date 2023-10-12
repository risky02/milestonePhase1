package cli

import (
	"fmt"
	"os"
	"os/exec"
	"project/config"
	"project/entity"
	"project/handler"
)

func Run() {
		fmt.Println("=====================================")
		fmt.Println("       Welcome to Sports Market      ")
		fmt.Println("=====================================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("9. Exit")
		fmt.Println("=====================================")

		input := choice()
		switch input {
		case 1:
			clearScreen()
			Login()
		case 2:
			clearScreen()
			register()
		case 9:
			clearScreen()
			fmt.Println("Terima kasih, sampai jumpa kembali")
			return
		default:
			fmt.Println("Wrong input")
		}
}

func clearScreen() {
	// clear screen windows
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func choice() int {
	// function input pilihan menu
	var input int
	fmt.Print("Choose menu: ")
	fmt.Scanln(&input)
	return input
}

func Login()  {
	var username, password string

	fmt.Println("===============LOGIN=================")
	fmt.Print("Masukkan username: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	user, isAuthenticated := handler.Login(username, password)
	if isAuthenticated {
		clearScreen()
		fmt.Println("Login berhasil! selamat datang,", user.Username)
		Menu()
	} else {
		clearScreen()
		fmt.Println("Login gagal!")
		Run()
	}

}

func register()  {
	var username, password string
	fmt.Println("============REGISTRASI===============")
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

func Menu()  {
	fmt.Println("=====================================")
	fmt.Println("       Toko E-commerce Olahraga      ")
	fmt.Println("=====================================")
	fmt.Println("1. Lihat Produk")
	fmt.Println("2. Tambahkan Produk ke Keranjang")
	fmt.Println("3. Lihat Keranjang Belanja")
	fmt.Println("4. Checkout")
	fmt.Println("9. Logout")
	fmt.Println("=====================================")

	input := choice()
	switch input {
	case 1:
		clearScreen()
		Category()
	case 2:
		fmt.Println("Tambahkan Produk")
	case 9:
		fmt.Println("Anda berhasil logout")
		clearScreen()
		Run()
	default:
		fmt.Println("Wrong input")
	}
	
}

func Category() {
    fmt.Println("========== 1. Lihat Produk ==========")
    fmt.Println("1. Lihat Semua Produk")
    fmt.Println("2. Lihat Produk Pakaian")
    fmt.Println("3. Lihat Produk Sepatu")
    fmt.Println("4. Lihat Produk Alat")
	fmt.Println("9. Back to Menu")

    var category string
    fmt.Printf("\nMasukkan Kategori: ")
    fmt.Scanln(&category)

    switch category {
    case "1":
        clearScreen()
        viewProducts("All")
    case "2":
        clearScreen()
        viewProducts("Pakaian") 
    case "3":
        clearScreen()
        viewProducts("Sepatu") 
    case "4":
        clearScreen()
        viewProducts("Alat")
	case "9":
		clearScreen()
		Menu()
    default:
		clearScreen()
        fmt.Println("Wrong input")
		Category()
    }
}

func viewProducts(category string) {
    db, err := config.InitDB()
    if err != nil {
        fmt.Println("Failed to connect to the database:", err)
        return
    }
    defer db.Close()

    var products []entity.Product
    err = nil

    if category == "All" {
        products, err = handler.ReadProducts(db)
    } else {
        products, err = handler.ReadCategories(db, category)
    }

    if err != nil {
        fmt.Println("Failed to retrieve products:", err)
        return
    }

    fmt.Println("List Products")
    fmt.Println("------------------------------------------------------------------")
    fmt.Println("| Product                     | Category          | Price        |")
    fmt.Println("------------------------------------------------------------------")

    for _, product := range products {
        fmt.Printf("| %-27s | %-17s | %12.2f |\n", product.ProductName, product.Category, product.Price)
    }

    fmt.Println("------------------------------------------------------------------")
    fmt.Print("\nPress ENTER to return to back...")
    fmt.Scanln()
    clearScreen()
    Category()
}

