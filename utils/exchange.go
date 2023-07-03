package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type OAuth2 struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

const BaseURL = "https://discord.com/api/"

func Exchange(code string) (OAuth2, int, string) {
	var oauth2 OAuth2
	body := EncodeParams(os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		code,
		os.Getenv("REDIRECT_URI"),
		os.Getenv("SCOPE"))
	res, err := http.Post(BaseURL+"oauth2/token", "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		log.Print("Error with the POST request", err)
	}
	var b []byte
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print("Error while trying to close body", err)
		}
	}(res.Body)
	if b, err = io.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}
	err = json.Unmarshal(b, &oauth2)
	if err != nil {
		log.Print("Error while parsing data to struct", err)
	}
	return oauth2, res.StatusCode, string(b)
}
