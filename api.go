package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	database Database
}

func (s *APIServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandler(s.handleCreateAccounts)).Methods("POST")
	router.HandleFunc("/account", makeHttpHandler(s.handleGetAccounts)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHttpHandler(s.handleGetAccountById)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHttpHandler(s.handleDeleteAccounts)).Methods("DELETE")
	router.HandleFunc("/transfer", makeHttpHandler(s.handleTransfer)).Methods("POST")


	log.Println("Starting server on port:" + s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func NewAPIServer(listenAddress string, db Database) *APIServer {
	return &APIServer{
		listenAddress: listenAddress, 
		database: db,
	}
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.database.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)

	idArg, err := strconv.Atoi(id["id"])
	
	if err != nil {
		return err
	}

	account, err := s.database.GetAccountById(idArg)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccounts(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateAccount)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	newAccount := NewAccount(req.FirstName, req.LastName)

	if err := s.database.CreateAccount(newAccount); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, newAccount)
}

func (s *APIServer) handleDeleteAccounts(w http.ResponseWriter, r *http.Request) error {
	
	id := mux.Vars(r)

	idArg, err := strconv.Atoi(id["id"])

	if err != nil {
		return err
	}

	if err := s.database.DeleteAccount(idArg); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHttpHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}