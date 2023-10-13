package handler

import (
	"project/config"
	"project/entity"
)

const (
    readCategories = `SELECT ProductName, Category, Price FROM products`
    readProducts   = `SELECT ProductName, Category, Price FROM products`
)

func ReadProducts() ([]entity.Product, error) {
	db, _ := config.InitDB()
    rows, err := db.Query(readProducts)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    products := []entity.Product{}
    for rows.Next() {
        var product entity.Product
        if err := rows.Scan(
            &product.ProductName,
            &product.Category,
            &product.Price,
        ); err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    return products, nil
}

func ReadCategories(category string) ([]entity.Product, error) {
	db, _ := config.InitDB()
    query := readCategories
    if category != "All" {
        query = readCategories + " WHERE Category = ?"
    }

    rows, err := db.Query(query, category) // Tambahkan parameter category
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    categories := []entity.Product{}
    for rows.Next() {
        var product entity.Product
        if err := rows.Scan(
            &product.ProductName,
            &product.Category,
            &product.Price,
        ); err != nil {
            return nil, err
        }
        categories = append(categories, product)
    }

    return categories, nil
}

