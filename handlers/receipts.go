package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt_app/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receipts = make(map[string]models.Receipt)

func CalculatePoints(receipt models.Receipt, total float64) int {
	points := 0

	reg := regexp.MustCompile("[a-zA-Z0-9]")
	points += len(reg.FindAllString(receipt.Retailer, -1))

	if total == math.Floor(total) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			continue
		}
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if descLength%3 == 0 {
			points += int(math.Ceil(price * 0.2))
		}
	}

	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil {
		if purchaseDate.Day()%2 == 1 {
			points += 6
		}
	}

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil {
		if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
			points += 10
		}
	}

	return points
}

func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	receipt.ID = id
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		http.Error(w, "Invalid total", http.StatusBadRequest)
		return
	}

	receipt.Points = CalculatePoints(receipt, total)
	receipts[id] = receipt

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["receiptId"]
	receipt, exists := receipts[receiptID]
	if !exists {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": receipt.Points})
}
