package handler

import (
	"database/config"
	"database/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func MainRouting() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", GetAll).Methods("GET")

	handler := cors.AllowAll().Handler(r)
	return handler
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	var user models.Customer
	result := config.DB.Find(&user, "user_id = ?", 3)
	if result.Error != nil {
		fmt.Println("Hata : ", result.Error)
		json.NewEncoder(w).Encode(result.Error)
	} else {
		json.NewEncoder(w).Encode(user)

	}
}
