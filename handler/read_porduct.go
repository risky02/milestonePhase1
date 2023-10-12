package handler

import (
	"database/sql"
	"project/entity"
)

const (
    readCategories = `SELECT ProductName, Category, Price FROM products`
    readProducts   = `SELECT ProductName, Category, Price FROM products`
)

func ReadProducts(db *sql.DB) ([]entity.Product, error) {
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

func ReadCategories(db *sql.DB, category string) ([]entity.Product, error) {
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

