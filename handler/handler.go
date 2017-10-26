package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stefaluc/cryptofolio-server/models"
)

type LoginParams struct {
	Username string
	Password string
}

func Login(c *gin.Context, in *LoginParams) (string, error) {
	user, err := models.GetUserFromLogin(in.Username)
	if err != nil {
		return "", err
	}

	// TODO: Check that hashing(in.Password) equals user.Password (already hashed)

	token, err := models.InsertToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

type SignUpParams struct {
	models.User
}

func SignUp(c *gin.Context, in *SignUpParams) error {
	// TODO: hashing(in.Password)
	_, err := models.InsertUser(&in.User)
	if err != nil {
		return err
	}
	return nil
}

type InsertBalanceParams struct {
	CryptocurrencyID int `json:"cryptocurrencyID"`
}

func InsertBalance(c *gin.Context, in *InsertBalanceParams) (*models.Balance, error) {
	// TODO: Get Token from Header
	t := models.STATIC_TOKEN
	user, err := models.GetUserFromToken(t)
	if err != nil {
		return nil, err
	}
	return models.InsertBalance(user, in.CryptocurrencyID)
}

type InsertTransactionParams struct {
	models.Transaction
}

func InsertTransaction(c *gin.Context, in *InsertTransactionParams) (*models.Transaction, error) {
	// TODO: Get Token from Header
	t := models.STATIC_TOKEN
	_, err := models.GetUserFromToken(t)
	if err != nil {
		return nil, err
	}
	// TODO: Check that the balance(in.Transaction.BalanceID) belongs to the user
	return models.InsertTransaction(&in.Transaction)
}

func GetUser(c *gin.Context) (*models.User, error) {
	// TODO: Get Token from Header
	t := models.STATIC_TOKEN
	return models.GetUserFromToken(t)
}

func GetBalances(c *gin.Context) ([]*models.Balance, error) {
	// TODO: Get Token from Header
	t := models.STATIC_TOKEN
	user, err := models.GetUserFromToken(t)
	if err != nil {
		return nil, err
	}
	return models.GetBalances(user)
}

type GetTransactionsParams struct {
	BalanceID int
}

func GetTransactions(c *gin.Context, in *GetTransactionsParams) ([]*models.Transaction, error) {
	// TODO: Get Token from Header
	t := models.STATIC_TOKEN
	_, err := models.GetUserFromToken(t)
	if err != nil {
		return nil, err
	}
	// TODO: Check that the balance(in.Transaction.BalanceID) belongs to the user
	return models.GetTransactions(in.BalanceID)
}

func GetCurrencies(c *gin.Context) ([]*models.Currency, error) {
	// TODO: Get Token from Header
	t := models.STATIC_TOKEN
	_, err := models.GetUserFromToken(t)
	if err != nil {
		return nil, err
	}
	return models.GetCurrencies()
}
