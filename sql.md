<!-- create database -->

CREATE DATABASE sports;

<!-- create table register -->

CREATE TABLE IF NOT EXISTS users (
Id INT AUTO_INCREMENT NOT NULL,
Username VARCHAR(50) UNIQUE NOT NULL,
Password VARCHAR(100) NOT NULL,
PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS products (
Id INT AUTO_INCREMENT NOT NULL,
ProductName VARCHAR(50) NOT NULL,
Category VARCHAR(30) NOT NULL,
Price DECIMAL (10, 2),
PRIMARY KEY (Id)
);

INSERT INTO `products` (ProductName, Category, Price)
VALUES
('Kaos Sepak Bola', 'Pakaian', '19.99'),
('Celana Renang Laki-Laki','Pakaian', 12.99),
('Sepatu Badminton', 'Sepatu', 49.99),
('Bola Basket', 'Alat', 9.99),
('Raket Tenis', 'Alat', 24.99);
