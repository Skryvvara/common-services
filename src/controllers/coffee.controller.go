package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func BrewCoffee(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	rng := rand.Intn(100 - 1)

	if rng > 1 {
		http.Error(w, http.StatusText(http.StatusTeapot), http.StatusTeapot)
		log.Println(fmt.Errorf("sadly there is no coffee to be found"))
		return
	}

	coffee := map[string]interface{}{
		"message": "I brewed fresh coffee, but sadly I can't send it over the internet, try the provided query instead",
		"query":   "https://google.com/search?q=coffee",
	}

	json.NewEncoder(w).Encode(coffee)
}
