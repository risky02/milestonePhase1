package handler

import (
	"context"
	"fmt"
	"project/config"
	"project/entity"
)

func AddCart(cart entity.Addcart) error {
	db, _ := config.InitDB()
	// Anda sebaiknya tidak membuat koneksi database baru di sini, tetapi gunakan yang dilewatkan
	ctx := context.Background()
	_, err := db.ExecContext(ctx, "INSERT INTO shopping_cart (UserId, ProductId, Quantity) VALUES (?, ?, ?)", cart.UserId, cart.ProductId, cart.Quantity)
	if err != nil {
		return err
	}
	fmt.Println("Success add item to shopping cart")
	return nil
}
