package handler

import (
	"Prisma/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Identify is an endpoint and execute the exchange function to get access tokens from Discord.
func Identify(w http.ResponseWriter, r *http.Request) {
	utils.UseCORS(&w, "*")
	w.Header().Set("Content-Type", "application/json")
	code := r.URL.Query().Get("code")
	value, status, reason := utils.Exchange(code)
	if status == http.StatusUnauthorized || status == http.StatusBadRequest {
		_, err := fmt.Fprint(w, reason)
		if err != nil {
			log.Fatal("Failed to print to screen", err)
		}
	} else {
		data, err := json.Marshal(value)
		if err != nil {
			log.Fatal("Error while parsing data to json", err)
		}
		fmt.Fprint(w, string(data))
	}
}
