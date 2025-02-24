# 📦 Trade Orders API

## 🚀 Overview
**Trade Orders API** is a backend service developed to handle **trade orders** in a financial system. Built with **Golang (Gin)** and **PostgreSQL**, the API allows users to create and retrieve trade orders efficiently. The application is containerized using **Docker**, deployed on **AWS EC2**, and integrated with **GitHub Actions CI/CD** for automated deployment.

---

## 🎯 Features
- **Order Management**: Users can submit and retrieve trade orders.
- **Database Storage**: Uses **PostgreSQL** for structured data management.
- **Containerization**: Runs inside **Docker** for portability.
- **CI/CD Integration**: Automated testing and deployment with **GitHub Actions**.
- **Cloud Deployment**: Hosted on **AWS EC2** for scalability.

---

## 🛠️ Tech Stack
- **Backend**: Golang (Gin Framework)
- **Database**: PostgreSQL
- **Containerization**: Docker
- **Deployment**: AWS EC2
- **CI/CD**: GitHub Actions

---

## 📥 Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/trade-orders-api.git
cd TradeOrder
```

### 2. Set Up PostgreSQL
#### If using **Docker**, run:
```bash
docker run -d --name postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=Cse@40668 \
  -e POSTGRES_DB=orders_db \
  -p 5432:5432 postgres
```
#### If using **local PostgreSQL**, update `main.go`:
```go
dsn := "host=localhost user=postgres password=Cse@40668 dbname=orders_db port=5432 sslmode=disable"
```

### 3. Install Dependencies & Run the Application
```bash
go mod tidy
go run main.go
```

### 4. Access the API
Open **Postman** or use `curl` to test:
```bash
curl -X POST http://localhost:8000/orders \
  -H "Content-Type: application/json" \
  -d '{
    "symbol": "AAPL",
    "price": 150.5,
    "quantity": 10,
    "order_type": "buy"
  }'
```

---

## 📦 Docker Setup

### 1. Build and Run the Docker Container
```bash
docker build -t orderservice .
docker run -p 8000:8000 orderservice
```

---

## 🚀 Deployment on AWS EC2

### 1. Connect to EC2
```bash
ssh -i your-key.pem ubuntu@your-ec2-ip
```

### 2. Deploy Docker Containers
```bash
docker pull your-dockerhub-username/orderservice:latest
docker stop orderservice || true
docker rm orderservice || true
docker run -d -p 8000:8000 --name orderservice your-dockerhub-username/orderservice
```

---

## 🔄 CI/CD Pipeline (GitHub Actions)

### Workflow Automation
✅ **Runs tests on PRs**  
✅ **Builds & pushes Docker image to Docker Hub**  
✅ **SSHs into EC2 & deploys the latest version**  

## 📌 API Documentation

### 1. Create a New Order
#### `POST /orders`
- **Request Example**:
```json
{
  "symbol": "AAPL",
  "price": 150.5,
  "quantity": 10,
  "order_type": "buy"
}
```
- **Response Example**:
```json
{
  "id": 1,
  "symbol": "AAPL",
  "price": 150.5,
  "quantity": 10,
  "order_type": "buy"
}
```

### 2. Retrieve All Orders
#### `GET /orders`
- **Response Example**:
```json
[
  {
    "id": 1,
    "symbol": "AAPL",
    "price": 150.5,
    "quantity": 10,
    "order_type": "buy"
  }
]
```

---

## 📚 Future Enhancements

- **🔒 Authentication & Authorization**: Secure API endpoints with JWT.
- **📊 Real-time Order Updates**: Implement WebSockets for live tracking.
- **☁️ Kubernetes Deployment**: Scale the application for high availability.
- **🧪 Unit & Integration Tests**: Improve reliability with automated tests.

---

## 🌎 Deployment
- **Currently Deployed On**: AWS EC2
- **Next Step**: Automate scaling with Kubernetes

---

## 📝 License
This project is open-source and available under the **MIT License**.

---
