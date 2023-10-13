package cli

import (
	"fmt"
	"os"
	"os/exec"
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
		clearScreen()
		Run()
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

func Login() {
    var username, password string

    fmt.Println("===============LOGIN=================")
    fmt.Print("Masukkan username: ")
    fmt.Scanln(&username)
    fmt.Print("Masukkan password: ")
    fmt.Scanln(&password)

    user, isAuthenticated := handler.Login(username, password)
    if isAuthenticated {
        clearScreen()
        fmt.Println("Login berhasil! Selamat datang,", user.Username)
        Menu(user.UserId)
    } else {
        clearScreen()
        fmt.Println("Login gagal!")
        Run()
    }
}

func register() {
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
        clearScreen()
        fmt.Println("Gagal Registrasi", dataRegist)
        Run()
    } else {
        clearScreen()
        fmt.Println("Berhasil registrasi")
        Run()
    }
}

func Menu(UserId int) {
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
        Category(UserId)
    case 2:
        clearScreen()
        AddToCart(UserId)
	case 3:
        clearScreen()
		viewShoppingCart(UserId)
	case 4:
        clearScreen()
        Checkout(UserId)
    case 9:
        fmt.Println("Anda berhasil logout")
        clearScreen()
        Run()
    default:
		clearScreen()
        fmt.Println("Wrong input")
		Menu(UserId)
    }
}

func Category(UserId int) {
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
        viewProducts("All", UserId)
    case "2":
        clearScreen()
        viewProducts("Pakaian", UserId)
    case "3":
        clearScreen()
        viewProducts("Sepatu", UserId)
    case "4":
        clearScreen()
        viewProducts("Alat", UserId)
    case "9":
        clearScreen()
        Menu(UserId)
    default:
        clearScreen()
        fmt.Println("Wrong input")
        Category(UserId)
    }
}

func viewProducts(category string, UserId int) {
    var products []entity.Product
    var err error

    if category == "All" {
        products, err = handler.ReadProducts()
    } else {
        products, err = handler.ReadCategories(category)
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
    Category(UserId)
}

func AddToCart(UserId int) {
    var productID int
    var quantity int

    fmt.Println("========== 2. Tambahkan Produk ke Keranjang ==========")
    fmt.Print("Masukkan Product ID: ")
    _, err := fmt.Scanf("%d\n", &productID)
    if err != nil {
		clearScreen()
        fmt.Println("ID Produk dan jumlah harus berupa angka.")
        AddToCart(UserId)
    }

    fmt.Print("Masukkan Jumlah: ")
    _, err = fmt.Scanf("%d\n", &quantity)
    if err != nil {
		clearScreen()
        AddToCart(UserId)
    }

	if quantity < 1 {
        clearScreen()
        fmt.Println("Jumlah minimum adalah 1. Produk akan ditambahkan dengan jumlah 1.")
        quantity = 1
    }

	cart := entity.Addcart{
		UserId:    UserId,
		ProductId: productID,
		Quantity:  quantity,
	}


    err = handler.AddCart(cart)
    if err != nil {
        fmt.Println("Gagal menambahkan produk ke keranjang belanja" )
    } else {
        fmt.Println("Produk berhasil ditambahkan ke keranjang belanja.")
    }

    fmt.Print("\nPress ENTER to return to the main menu...")
    fmt.Scanln()
    clearScreen()
    Menu(UserId)
}

func viewShoppingCart1(UserId int)  {
	listcart, products, err := handler.Whistlist2(UserId)

    if err != nil {
        fmt.Println("Failed to retrieve whistlist:", err)
        return
    }
	
	priceTotal := calculatePrice(listcart, products)

    fmt.Println("Shopping Cart Contents")
    fmt.Println("---------------------------------------------------------")
    fmt.Println("| Product Name              | Quantity   | Price Total  |")
    fmt.Println("---------------------------------------------------------")

    for i, cart := range listcart {
        // fmt.Printf("| %-25s | %-10d |\n", cart.ProductName, cart.Quantity)
		fmt.Printf("| %-25s | %-10d | $%-10.2f  |\n", cart.ProductName, cart.Quantity, priceTotal[i])
    }

    fmt.Println("---------------------------------------------------------")
}

func viewShoppingCart(UserId int) {
    // Menggunakan handler.Whistlist untuk mendapatkan data whistlist
    viewShoppingCart1(UserId)
    fmt.Print("\nPress ENTER to return to the main menu...")
    fmt.Scanln()
    clearScreen()
    Menu(UserId)
}

func calculatePrice(listcart []entity.ListCart, products []entity.Product) map[int]float64 {
    priceTotal := make(map[int]float64)

    for i, cart := range listcart {
        priceTotal[i] = float64(cart.Quantity) * products[i].Price
    }

    return priceTotal
}

func Checkout(UserId int) {
    viewShoppingCart(UserId)
    fmt.Printf("Total Belanja: $")
	var input string
	fmt.Println("Apakah anda ingin melakukan checkout? (y/n) ", input)
	if input == "y" || input == "Y" {
		fmt.Println("Tersimpan di database")
		// handler.AddCheckout()
	} else if input == "n" || input == "N" {
		clearScreen()
		Menu(UserId)
	} else {
		clearScreen()
		fmt.Println("Wrong Input")
		Checkout(UserId)
	}
}

func calculateTotalAmount(listcart []entity.ListCart, products []entity.Product) float64 {
    totalAmount := 0.0

    for i, cart := range listcart {
        totalAmount += float64(cart.Quantity) * products[i].Price
    }

    return totalAmount
}

