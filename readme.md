# 🏦 Wallet API (Golang + Echo + SQLX + Clean Architecture)

A simple digital wallet REST API with support for:
- ✅ User Register/Login (JWT auth)
- 💰 Withdraw & Deposit
- 📊 Balance Inquiry
- 🧾 Transaction History
- 🔐 JWT Middleware
- 🐳 Docker-ready
- 🧪 MySQL Seeder & Migration
- 📄 Swagger Docs

---

## 🚀 Tech Stack

- Go (1.21+)
- Echo Web Framework
- SQLX for MySQL access
- Clean Architecture (Handler / Usecase / Repository / Domain)
- JWT (`github.com/golang-jwt/jwt/v5`)
- Swagger (OpenAPI 3.0)

---

## ⚙️ Setup

### 1. Clone & Build

```bash
git clone <repo>
cd wallet-api
go mod tidy
go run main.go
```

### 2. Create `.env`

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=wallet_db
JWT_SECRET=my_secret_key
```

### 3. Run Migration + Seeder

```bash
go run migration/migrate.go
```

### 4. Run with Docker

```bash
docker build -t wallet-api .
docker run -p 8080:8080 --env-file .env wallet-api
```

---

## 🔐 API Endpoints

| Endpoint         | Method | Auth | Description            |
|------------------|--------|------|------------------------|
| `/register`      | POST   | ❌    | Register new user      |
| `/login`         | POST   | ❌    | Login and get JWT      |
| `/api/balance`   | GET    | ✅    | Get wallet balance     |
| `/api/deposit`   | POST   | ✅    | Add funds              |
| `/api/withdraw`  | POST   | ✅    | Withdraw funds         |
| `/api/history`   | GET    | ✅    | List transaction logs  |

---

## 📄 Swagger Docs

1. Serve via local:

```bash
cd docs
python3 -m http.server
```

2. Open browser: [http://localhost:8000](http://localhost:8000)

---

## 🧪 Test Users (from Seeder)

| Username | Password   |
|----------|------------|
| alice    | secret123  |

---

## 📁 Structure

```bash
wallet-api/
├── main.go
├── .env
├── Dockerfile
├── migration/
├── config/
├── docs/
├── internal/
│   ├── domain/
│   ├── handler/
│   ├── repository/
│   └── usecase/
```

---

MIT License
