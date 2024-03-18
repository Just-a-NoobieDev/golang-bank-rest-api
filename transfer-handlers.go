package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

	req := new(CreateTransfer)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid request body"})
	}

	newTransfer := NewTransfer(req.From, req.To, req.Amount)

	if err := s.database.CreateTransfer(newTransfer); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Something went wrong creating your transfer. Please Try again later." + err.Error()})
	}

	return WriteJSON(w, http.StatusCreated, newTransfer)
}

func (s *APIServer) handleGetTransfers(w http.ResponseWriter, r *http.Request) error {
	transfers, err := s.database.GetTransfers()
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Something went wrong. Please try again later."})
	}
	return WriteJSON(w, http.StatusOK, transfers)
}

func (s *APIServer) handleGetTransferByFrom(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)

	idArg, err := strconv.Atoi(id["id"])

	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid id"})
	}

	transfer, err := s.database.GetTransferByFrom(idArg)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Transfer not found."})
	}
	return WriteJSON(w, http.StatusOK, transfer)
}

func (s *APIServer) handleGetTransferByTo(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)

	idArg, err := strconv.Atoi(id["id"])

	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid id"})
	}

	transfer, err := s.database.GetTransferByTo(idArg)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Transfer not found."})
	}
	return WriteJSON(w, http.StatusOK, transfer)
}