package handler

import (
	"project/config"
	"project/entity"
)

func Whistlist(UserId int) ([]entity.ListCart, error) {
    db, _ := config.InitDB()

    rows, err := db.Query(`
        SELECT p.ProductName, sc.Quantity
        FROM shopping_cart sc
        JOIN products p ON sc.ProductId = p.ProductId
        WHERE sc.UserId = ?;
    `, UserId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var listcart []entity.ListCart
    for rows.Next() {
        var cart entity.ListCart
        err := rows.Scan(&cart.ProductName, &cart.Quantity)
        if err != nil {
            return nil, err
        }
        listcart = append(listcart, cart)
    }

    return listcart, nil
}


func Whistlist2(UserId int) ([]entity.ListCart, []entity.Product, error) {
    db, _ := config.InitDB()

    rows, err := db.Query(`
        SELECT p.ProductName, sc.Quantity, p.Price
        FROM shopping_cart sc
        JOIN products p ON sc.ProductId = p.ProductId
        WHERE sc.UserId = ?;
    `, UserId)
	if err != nil {
        return nil, nil, err
    }
    defer rows.Close()


    var listcart []entity.ListCart
    var products []entity.Product
    
    for rows.Next() {
        for rows.Next() {
			var cart entity.ListCart
			var product entity.Product
			err := rows.Scan(&cart.ProductName, &cart.Quantity, &product.Price)
			if err != nil {
				return nil, nil, err
			}
			listcart = append(listcart, cart)
			products = append(products, product)
		}
	}
	return listcart, products, nil
}