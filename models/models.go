package models

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/stefaluc/cryptofolio-server/database"
)

type User struct {
	ID                  int    `json:"-"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	FavouriteCurrencyID int    `json:"favouriteCurrency"`
}

type Currency struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logoURL"`
}

type Balance struct {
	ID         int `json:"id"`
	UserID     int `json:"userID"`
	CurrencyID int `json:"currency_id"`
}

type Transaction struct {
	ID        int    `json:"id"`
	BalanceID int    `json:"balanceID"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	Date      string `json:"date"`
}

type Token struct {
	Token  string
	UserID int
}

var STATIC_TOKEN = &Token{
	Token:  "XXXXXXXXXXXXX",
	UserID: 1,
}
var STATIC_USER = &User{
	ID:                  1,
	Username:            "BGdu59",
	Password:            "secretpassword",
	FirstName:           "Jean",
	LastName:            "Dupont",
	FavouriteCurrencyID: 1,
}
var STATIC_CURRENCY = &Currency{
	ID:      1,
	Name:    "Bitcoin",
	LogoURL: "https://bitcoin.org/img/icons/opengraph.png",
}
var STATIC_BALANCE = &Balance{
	ID:         1,
	UserID:     1,
	CurrencyID: 1,
}
var STATIC_TRANSACTION = &Transaction{
	ID:        1,
	BalanceID: 1,
	Quantity:  1,
	Price:     10,
	Date:      "TODO",
}

func GetToken(username string, password string) (string, error) {
	return STATIC_TOKEN.Token, nil
}

func InsertToken(u *User) (string, error) {
	// TODO: Create Token in DB
	return STATIC_TOKEN.Token, nil
}

func GetUserFromLogin(username string) (*User, error) {
	// Get User with Token t in DB
	row := database.DBConn.QueryRow("SELECT * FROM \"user\" WHERE username=$1", username)

	var id int
	var usernameDb string
	var password string
	var first_name string
	var last_name string
	var favourite_currency_id int
	err := row.Scan(&id, &usernameDb, &password, &first_name, &last_name, &favourite_currency_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("id | username | password | firstname | lastname | favouritecurrency")
	fmt.Printf("%3v | %8v | %6v | %6v | %6v | %3v\n", id, usernameDb, password, first_name, last_name, favourite_currency_id)

	return &User{
		ID:                  id,
		Username:            username,
		Password:            password,
		FirstName:           first_name,
		LastName:            last_name,
		FavouriteCurrencyID: favourite_currency_id,
	}, nil
}

func GetUserFromToken(t *Token) (*User, error) {
	// TODO: Get User with Token t in DB
	return STATIC_USER, nil
}

func InsertUser(u *User) (*User, error) {
	// query for preexisting username
	var username string
	err := database.DBConn.QueryRow("SELECT username FROM user WHERE username=$1").Scan(&username)
	switch {
	// valid new user, hash password and insert into db
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		_, err = database.DBConn.Exec(
			"INSERT INTO \"user\"(id, username, password, first_name, last_name, favourite_currency_id) VALUES($1,$2,$3,$4,$5,$6) returning id;",
			u.ID, u.Username, hashedPassword, u.FirstName, u.LastName, u.FavouriteCurrencyID)
		if err != nil {
			return nil, err
		}
	case err != nil:
		return nil, err
	// username already exists in db
	default:
		// TODO: handle
	}

	return nil, nil
}

func InsertBalance(u *User, crypto int) (*Balance, error) {
	return STATIC_BALANCE, nil
}

func InsertTransaction(t *Transaction) (*Transaction, error) {
	// TODO: Create transaction in DB
	return STATIC_TRANSACTION, nil
}

func GetBalances(u *User) ([]*Balance, error) {
	// TODO
	return []*Balance{STATIC_BALANCE}, nil
}

func GetTransactions(balanceID int) ([]*Transaction, error) {
	// TODO
	return []*Transaction{STATIC_TRANSACTION}, nil
}

func GetCurrencies() ([]*Currency, error) {
	// TODO
	return []*Currency{STATIC_CURRENCY}, nil
}
