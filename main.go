package main

import (
	"fmt"
	"log"
	"net/http"
	"receipt_app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", handlers.CreateReceipt).Methods("POST")
	r.HandleFunc("/receipts/{receiptId}/points", handlers.GetPoints).Methods("GET")

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
