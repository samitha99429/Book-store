package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strings"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"book-store/models"
)

func GetBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(utils.Books)
}

func Create