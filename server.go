package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	database Database
}

func (s *APIServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/accounts", makeHttpHandler(s.handleCreateAccounts)).Methods("POST")
	router.HandleFunc("/api/v1/accounts", makeHttpHandler(s.handleGetAccounts)).Methods("GET")
	router.HandleFunc("/api/v1/accounts/{id}", makeHttpHandler(s.handleGetAccountById)).Methods("GET")
	router.HandleFunc("/api/v1/accounts/{id}", makeHttpHandler(s.handleDeleteAccounts)).Methods("DELETE")
	router.HandleFunc("/api/v1/transfer", makeHttpHandler(s.handleTransfer)).Methods("POST")
	router.HandleFunc("/api/v1/transfers", makeHttpHandler(s.handleGetTransfers)).Methods("GET")
	router.HandleFunc("/api/v1/transfers-from/{id}", makeHttpHandler(s.handleGetTransferByFrom)).Methods("GET")
	router.HandleFunc("/api/v1/transfers-to/{id}", makeHttpHandler(s.handleGetTransferByTo)).Methods("GET")

	log.Println("Starting server on port:" + s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func NewAPIServer(listenAddress string, db Database) *APIServer {
	return &APIServer{
		listenAddress: listenAddress, 
		database: db,
	}
}





