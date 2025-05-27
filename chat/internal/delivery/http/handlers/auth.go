package handlers

import (
	"net/http"
	
	"github.com/gorilla/mux"
)

func InitAuth(r *mux.Router) {
	r.HandleFunc("/login", Login).Methods("GET")
	r.HandleFunc("/register", Register).Methods("GET")
}

func Login(w http.ResponseWriter, r *http.Request) {}

func Register(w http.ResponseWriter, r *http.Request) {}