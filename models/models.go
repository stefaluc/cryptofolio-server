package models

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
	// TODO: Get User with Token t in DB
	return STATIC_USER, nil
}

func GetUserFromToken(t *Token) (*User, error) {
	// TODO: Get User with Token t in DB
	return STATIC_USER, nil
}

func InsertUser(u *User) (*User, error) {
	// TODO: Create User in DB
	return STATIC_USER, nil
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
