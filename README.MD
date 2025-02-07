# Receipt Processor API

This is a Go-based API for processing receipts and calculating reward points based on predefined rules. The application uses `gorilla/mux` for routing and provides two main API routes:

- `POST /receipts/process` - Accepts a receipt and returns an assigned ID.
- `GET /receipts/{receiptId}/points` - Retrieves the points associated with a receipt.

---

## **Getting Started**

You can run this application in two ways:
1. **Without Docker** (Running directly on your system)
2. **With Docker** (Running inside a container)

---

## **1. Running Locally Without Docker**

### **Prerequisites**
Ensure you have the following installed on your system:
- [Go 1.23.6](https://go.dev/dl/) or later
- `curl` (for testing API requests)

### **Clone the Repository**
```bash
git clone https://github.com/your-username/receipt-processor.git
cd receipt-processor
```

### **Install Dependencies**
```
go mod tidy
```

### **Run the Application**
```
go run main.go
```

### **Test the API using Curl**
```
curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '{
          "retailer": "Walmart",
          "purchaseDate": "2024-02-07",
          "purchaseTime": "15:30",
          "total": "35.50",
          "items": [
              { "shortDescription": "Milk", "price": "3.50" },
              { "shortDescription": "Bread", "price": "2.00" }
          ]
      }'
```
This will return a JSON response with an assigned receipt ID.

### **Retrieve receipt points**
```
curl http://localhost:8080/receipts/{receiptId}/points
```
Replace {receiptId} with the actual ID from the previous response.

## **2. Running With Docker**

### **Build the Docker Image**

```
docker build -t receipt-processor .
```

### **Run the Container**

```
docker run -p 8080:8080 receipt-processor
```

### **Test the API inside Docker**
```
curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '{
          "retailer": "Target",
          "purchaseDate": "2024-02-08",
          "purchaseTime": "14:45",
          "total": "42.00",
          "items": [
              { "shortDescription": "Apples", "price": "4.00" },
              { "shortDescription": "Bananas", "price": "3.00" }
          ]
      }'
```
This will return a JSON response with an assigned receipt ID.

### **Retrieve receipt points**
```
curl http://localhost:8080/receipts/{receiptId}/points
```
