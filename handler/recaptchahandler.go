package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	recaptchaSecret = "6LfUazwUAAAAAN0bXeI8-3I3zISPpzg40i6fqrWA"
)

type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	HostName    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

func VerifyRecaptcha(userResponse string) error {
	// POST to google recaptcha api with secret+usertoken
	postUrl := "https://www.google.com/recaptcha/api/siteverify"
	vals := url.Values{}
	vals.Set("secret", recaptchaSecret)
	vals.Add("response", userResponse)

	res, err := http.PostForm(postUrl, vals)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// decode response
	decoder := json.NewDecoder(res.Body)

	resJson := &RecaptchaResponse{}
	err = decoder.Decode(&resJson)
	if err != nil {
		return err
	}
	fmt.Println(resJson)

	if !resJson.Success {
		return errors.New("Invalid success response from Google")
	}
	return nil
}
