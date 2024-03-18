package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.database.GetAccounts()
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Something went wrong. Please try again later."})
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)

	idArg, err := strconv.Atoi(id["id"])

	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid id"})
	}

	account, err := s.database.GetAccountById(idArg)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "User not found."})
	}
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccounts(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateAccount)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid request body. Please provide a valid first and last name."})
	}

	newAccount := NewAccount(req.FirstName, req.LastName)

	if err := s.database.CreateAccount(newAccount); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Something went wrong creating your account. Please Try again later."})
	}

	return WriteJSON(w, http.StatusCreated, newAccount)
}

func (s *APIServer) handleDeleteAccounts(w http.ResponseWriter, r *http.Request) error {

	id := mux.Vars(r)

	idArg, err := strconv.Atoi(id["id"])

	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid id"})
	}

	if err := s.database.DeleteAccount(idArg); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Something went wrong deleting your account. Please Try again later."})
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"status": "success"})
}