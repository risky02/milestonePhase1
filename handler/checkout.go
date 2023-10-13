package handler

import (
	"project/config"
	"project/entity"
	"time"
)

// Checkout digunakan untuk menambahkan checkout baru ke database.
func AddCheckout(checkout entity.Checkout) error {
	db, _ := config.InitDB()
    // Menyimpan data checkout ke database.
    _, err := db.Exec("INSERT INTO checkout (UserID, TotalAmount, CheckoutDate) VALUES (?, ?, ?)",
        checkout.UserID, checkout.TotalAmount, time.Now())
    if err != nil {
        return err
    }

    return nil
}


func GetProductsForShoppingCart(listcart []entity.ListCart) ([]entity.Product, error) {
    db, err := config.InitDB()
    if err != nil {
        return nil, err
    }

    // Create a slice to store product information
    var products []entity.Product

    for _, cart := range listcart {
        var product entity.Product

        // Retrieve product information for each item in the shopping cart
        err := db.QueryRow("SELECT ProductID, ProductName, Price FROM products WHERE ProductID = ?", cart.ProductId).Scan(&product.ProductId, &product.ProductName, &product.Price)
        if err != nil {
            return nil, err
        }

        products = append(products, product)
    }

    return products, nil
}

