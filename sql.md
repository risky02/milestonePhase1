<!-- create database -->

CREATE DATABASE sports;

<!-- create table register -->

CREATE TABLE IF NOT EXISTS users (
Id INT AUTO_INCREMENT NOT NULL,
Username VARCHAR(50) UNIQUE NOT NULL,
Password VARCHAR(100) NOT NULL,
PRIMARY KEY (Id)
);
