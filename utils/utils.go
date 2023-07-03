package utils

import (
	"net/http"
	"net/url"
)

/*
EncodeParams formats and add the needed params by discord, it is used for OAuth2 token exchanges.
*/
func EncodeParams(clientId, clientSecret, code, redirectURI, scope string) string {
	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", redirectURI)
	params.Add("scope", scope)
	return params.Encode()
}

/*
UseCORS add CORS header, you can pass as a string the allowed origin.
The * wildcard will allow any origin.
*/
func UseCORS(w *http.ResponseWriter, origin string) {
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
}
