package models

import (
	"database/sql"
	// "fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/stefaluc/cryptofolio-server/database"
)

type User struct {
	ID                  int    `json:"-"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	FavouriteCurrencyID int    `json:"favouriteCurrencyID"`
}

type Currency struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logoURL"`
}

type Balance struct {
	ID         int     `json:"id"`
	UserID     int     `json:"userID"`
	CurrencyID int     `json:"currencyID"`
	Quantity   float32 `json:"quantity"`
}

type Transaction struct {
	ID        int     `json:"id"`
	BalanceID int     `json:"balanceID"`
	Quantity  float32 `json:"quantity"`
	Price     float32 `json:"price"`
	Date      string  `json:"date"`
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

func GetUserFromUsername(username string) (*User, error) {
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
	// fmt.Println("id | username | password | firstname | lastname | favouritecurrency")
	// fmt.Printf("%3v | %8v | %6v | %6v | %6v | %3v\n", id, usernameDb, password, first_name, last_name, favourite_currency_id)

	return &User{
		ID:                  id,
		Username:            username,
		Password:            password,
		FirstName:           first_name,
		LastName:            last_name,
		FavouriteCurrencyID: favourite_currency_id,
	}, nil
}

func InsertUser(u *User) (*User, error) {
	// query for preexisting username
	var username string
	err := database.DBConn.QueryRow("SELECT username FROM \"user\" WHERE username=$1", u.Username).Scan(&username)

	var id int
	var user *User
	switch {
	// valid new user, hash password and insert into db
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		err = database.DBConn.QueryRow(
			"INSERT INTO \"user\"(username, password, first_name, last_name, favourite_currency_id) VALUES($1,$2,$3,$4,$5) returning id;",
			u.Username, string(hashedPassword), u.FirstName, u.LastName, u.FavouriteCurrencyID).Scan(&id)
		if err != nil {
			return nil, err
		}

		user = &User{
			ID:                  id,
			Username:            username,
			Password:            string(hashedPassword),
			FirstName:           u.FirstName,
			LastName:            u.LastName,
			FavouriteCurrencyID: u.FavouriteCurrencyID,
		}
	case err != nil:
		return nil, err
	// username already exists in db
	default:
		// TODO: handle
	}

	return user, nil
}

func InsertBalance(u *User, crypto int, quantity float32) (*Balance, error) {
	var id int
	err := database.DBConn.QueryRow(
		"INSERT INTO \"balance\"(user_id, currency_id, quantity) VALUES($1,$2,$3) returning id;",
		u.ID, crypto, quantity).Scan(&id)
	if err != nil {
		return nil, err
	}

	balance := &Balance{
		ID:         id,
		UserID:     u.ID,
		CurrencyID: crypto,
		Quantity:   quantity,
	}
	return balance, nil
}

func InsertTransaction(t *Transaction) (*Transaction, error) {
	var id int
	err := database.DBConn.QueryRow(
		"INSERT INTO \"transaction\"(balance_id, quantity, price, date) VALUES($1,$2,$3,$4) returning id;",
		t.BalanceID, t.Quantity, t.Price, t.Date).Scan(&id)
	if err != nil {
		return nil, err
	}

	transaction := &Transaction{
		ID:        id,
		BalanceID: t.BalanceID,
		Quantity:  t.Quantity,
		Price:     t.Price,
		Date:      t.Date,
	}
	return transaction, nil
}

func GetBalances(u *User) ([]*Balance, error) {
	rows, err := database.DBConn.Query("SELECT * FROM \"balance\" WHERE user_id=$1", u.ID)
	if err != nil {
		return nil, err
	}

	var balances []*Balance
	for rows.Next() {
		var id int
		var userID int
		var currencyID int
		var quantity float32
		err := rows.Scan(&id, &userID, &currencyID, &quantity)
		if err != nil {
			return nil, err
		}

		balance := &Balance{
			ID:         id,
			UserID:     u.ID,
			CurrencyID: currencyID,
			Quantity:   quantity,
		}
		balances = append(balances, balance)
	}

	return balances, nil
}

func GetTransactions(balanceID int) ([]*Transaction, error) {
	rows, err := database.DBConn.Query("SELECT * FROM \"transaction\" WHERE balance_id=$1", balanceID)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	for rows.Next() {
		var id int
		var balanceID int
		var quantity float32
		var price float32
		var date string
		err := rows.Scan(&id, &balanceID, &quantity, &price, &date)
		if err != nil {
			return nil, err
		}

		transaction := &Transaction{
			ID:        id,
			BalanceID: balanceID,
			Quantity:  quantity,
			Price:     price,
			Date:      date,
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func GetCurrencies() ([]*Currency, error) {
	rows, err := database.DBConn.Query("SELECT * From \"currency\"")
	if err != nil {
		return nil, err
	}

	var currencies []*Currency
	for rows.Next() {
		var id int
		var name string
		var logoURL string
		err := rows.Scan(&id, &name, &logoURL)
		if err != nil {
			return nil, err
		}

		currency := &Currency{
			ID:      id,
			Name:    name,
			LogoURL: logoURL,
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}
