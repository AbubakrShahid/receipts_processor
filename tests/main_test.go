package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/receipts/process/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"id": "123e4567-e89b-12d3-a456-426614174000"})
			return
		}

		if r.Method == http.MethodGet && len(r.URL.Path) > 10 && r.URL.Path[:10] == "/receipts/" {
			var receiptID string
			fmt.Sscanf(r.URL.Path, "/receipts/%s/points", &receiptID)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]int{"points": 150})
			return
		}

		http.NotFound(w, r)
	}))
	defer testServer.Close()

	postURL := testServer.URL + "/receipts/process/"
	reqBody := []byte(`{}`)
	req, err := http.NewRequest(http.MethodPost, postURL, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading POST response:", err)
		return
	}

	var postResponse map[string]string
	if err := json.Unmarshal(body, &postResponse); err != nil {
		fmt.Println("Error parsing POST response:", err)
		fmt.Println("Response Body:", string(body))
		return
	}

	receiptID, exists := postResponse["id"]
	if !exists {
		fmt.Println("Error: ID not found in POST response")
		return
	}

	fmt.Println("Received Receipt ID:", receiptID)

	getURL := fmt.Sprintf("%s/receipts/%s/points", testServer.URL, receiptID)
	req, err = http.NewRequest(http.MethodGet, getURL, nil)
	if err != nil {
		fmt.Println("Error creating GET request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading GET response:", err)
		return
	}

	var getResponse map[string]int
	if err := json.Unmarshal(body, &getResponse); err != nil {
		fmt.Println("Error parsing GET response:", err)
		fmt.Println("Response Body:", string(body))
		return
	}

	points, exists := getResponse["points"]
	if !exists {
		fmt.Println("Error: Points not found in GET response")
		return
	}

	fmt.Println("Received Points:", points)
}
