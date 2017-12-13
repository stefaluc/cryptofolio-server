package handler

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/stefaluc/cryptofolio-server/models"
)

type LoginParams struct {
	Username string
	Password string
}

func Login(c *gin.Context, in *LoginParams) (string, error) {
	userDb, err := models.GetUserFromUsername(in.Username)
	if err != nil {
		return "", err
	}

	// check if password is same as db hashed password
	err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(in.Password))
	if err != nil {
		return "", err
	}

	// create auth token and send to client
	token, err := CreateToken(userDb)
	if err != nil {
		return "", err
	}

	return token, nil
}

type SignUpParams struct {
	models.User
	GRecaptchaResponse string `json:"gRecaptchaResponse"`
}

func SignUp(c *gin.Context, in *SignUpParams) error {
	// verify correct Google Recaptcha response
	err := VerifyRecaptcha(in.GRecaptchaResponse)
	if err != nil {
		return err
	}

	_, err = models.InsertUser(&in.User)
	if err != nil {
		return err
	}
	return nil
}

type InsertBalanceParams struct {
	CryptocurrencyID int     `json:"cryptocurrencyID"`
	Quantity         float32 `json:"quantity"`
}

func InsertBalance(c *gin.Context, in *InsertBalanceParams) (*models.Balance, error) {
	// check for valid token
	token := c.Query("token")
	username, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	user, err := models.GetUserFromUsername(username)
	if err != nil {
		return nil, err
	}

	return models.InsertBalance(user, in.CryptocurrencyID, in.Quantity)
}

type InsertTransactionParams struct {
	models.Transaction
}

func InsertTransaction(c *gin.Context, in *InsertTransactionParams) (*models.Transaction, error) {
	// check for valid token
	token := c.Query("token")
	_, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	balanceID, err := strconv.Atoi(c.Param("balance"))
	in.BalanceID = balanceID

	// TODO: Check that the balance(in.Transaction.BalanceID) belongs to the user
	return models.InsertTransaction(&in.Transaction)
}

func GetUser(c *gin.Context) (*models.User, error) {
	// check for valid token
	token := c.Query("token")
	username, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	user, err := models.GetUserFromUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetBalances(c *gin.Context) ([]*models.Balance, error) {
	// check for valid token
	token := c.Query("token")
	username, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	user, err := models.GetUserFromUsername(username)
	if err != nil {
		return nil, err
	}

	return models.GetBalances(user)
}

type GetTransactionsParams struct {
	BalanceID int
}

func GetTransactions(c *gin.Context, in *GetTransactionsParams) ([]*models.Transaction, error) {
	// check for valid token
	token := c.Query("token")
	_, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	balanceID, err := strconv.Atoi(c.Query("balanceID"))
	in.BalanceID = balanceID

	// TODO: Check that the balance(in.Transaction.BalanceID) belongs to the user
	return models.GetTransactions(in.BalanceID)
}

func GetCurrencies(c *gin.Context) ([]*models.Currency, error) {
	// check for valid token
	token := c.Query("token")
	_, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	return models.GetCurrencies()
}
