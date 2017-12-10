CREATE TABLE "user" (
    id serial NOT NULL PRIMARY KEY,
    username varchar(100) NOT NULL,
    password varchar(500) NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    favourite_currency_id INT NOT NULL
);

CREATE TABLE "balance" (
    id serial NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    currency_id INT NOT NULL,
    quantity REAL NOT NULL
);

CREATE TABLE "transaction" (
    id serial NOT NULL PRIMARY KEY,
    balance_id INT NOT NULL,
    quantity REAL NOT NULL,
    price REAL NOT NULL,
    date TIMESTAMP NOT NULL
);

CREATE TABLE "currency" (
    id serial NOT NULL PRIMARY KEY,
    name varchar(200) NOT NULL,
    logo_url varchar(500) NOT NULL
);
