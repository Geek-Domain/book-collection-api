# 📚 Book Collection API

A secure and RESTful backend API built with **Go**, **Gin**, and **MongoDB**, allowing users to register, log in, and manage a personal collection of books. Deployed on **Render**, with full JWT-based authentication and MongoDB Atlas for data storage.

---

## 🚀 Live Demo

🌐 **Base URL**: [https://geekbooks-com.onrender.com](https://geekbooks-com.onrender.com)

---

## ✅ Features

* User Registration & Login with password hashing
* JWT Authentication Middleware
* CRUD operations for books
* MongoDB Atlas integration
* Full REST API structure
* Secure environment variable handling
* Deployed live using Render

---

## 🧪 API Endpoints

### 🔐 Authentication

| Method | Endpoint    | Description        |
| ------ | ----------- | ------------------ |
| POST   | `/register` | Register a user    |
| POST   | `/login`    | Log in and get JWT |

### 📚 Book Management (Protected)

| Method | Endpoint     | Description           |
| ------ | ------------ | --------------------- |
| POST   | `/books`     | Create a new book     |
| GET    | `/books`     | List all user's books |
| GET    | `/books/:id` | Get book by ID        |
| PUT    | `/books/:id` | Update book by ID     |
| DELETE | `/books/:id` | Delete book by ID     |

> 🛡️ All `/books` routes require an `Authorization` header:
> `Bearer <your_jwt_token>`

---

## 🛠 Tech Stack

* **Language**: Go (Golang)
* **Framework**: Gin
* **Database**: MongoDB Atlas
* **Auth**: JWT (github.com/golang-jwt/jwt)
* **Hashing**: bcrypt
* **Deployment**: Render
* **Version Control**: Git + GitHub

---

## ⚙️ Run Locally

### 1. Clone the Repository

```bash
git clone https://github.com/Geek-Domain/book-collection-api.git
cd book-collection-api
```

### 2. Set up Environment Variables

Create a `.env` file:

```
MONGODB_URI=your_mongodb_uri
JWT_SECRET=your_jwt_secret
```

### 3. Run the API

```bash
go run main.go
```

The server will start at `http://localhost:8080`

---

## 📦 Project Structure

```
book-collection-api/
├── controllers/       # Auth and Book handlers
├── routes/            # Route registration
├── middleware/        # JWT middleware
├── models/            # User and Book structs
├── utils/             # JWT utilities
├── config/            # DB connection logic
├── main.go            # App entry point
├── go.mod / go.sum    # Dependencies
└── .env               # Environment config
```

---

## 📜 License

MIT License

---

## ✍️ Author

Built by **Geek Domain**
Feel free to reach out or contribute!
