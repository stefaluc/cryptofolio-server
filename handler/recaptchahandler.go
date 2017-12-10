package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	recaptchaSecret = "6LfUazwUAAAAAN0bXeI8-3I3zISPpzg40i6fqrWA"
)

type RecaptchaRequest struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

type RecaptchaResponse struct {
	Success     bool   `json:"success"`
	ChallengeTs string `json:"challenge_ts"`
	HostName    string `json:"hostname"`
}

func VerifyRecaptcha(userResponse string) error {
	url := "https://www.google.com/recaptcha/api/siteverify"

	req := &RecaptchaRequest{
		recaptchaSecret,
		userResponse,
	}
	reqJson, _ := json.Marshal(req)

	res, err := http.Post(url, "authentication/json", bytes.NewBuffer(reqJson))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	resJson := &RecaptchaResponse{}
	err = decoder.Decode(&resJson)
	if err != nil {
		return err
	}
	fmt.Println(resJson)
	fmt.Println(resJson.Success)

	if !resJson.Success {
		return errors.New("Invalid success response from Google")
	}
	return nil
}
