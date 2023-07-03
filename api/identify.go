package handler

import (
	"Prisma/utils"
	c "Prisma/utils/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Identify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := r.URL.Query().Get("code")
	value, status := utils.Exchange(c.Configuration, code)
	if status == http.StatusBadRequest {
		_, err := fmt.Fprint(w, "Invalid Code")
		if err != nil {
			return
		}
	} else {
		data, err := json.Marshal(value)
		if err != nil {
			log.Print("Error while parsing data to json", err)
		}
		fmt.Fprint(w, string(data))
	}
}
