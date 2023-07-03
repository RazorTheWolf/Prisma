package utils

import "net/url"

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
