# 📝 Blog Management System - Golang REST API

Sebuah project mini RESTful API dengan fitur user registration, login, content creation, category management, dan role-based access control. Dibangun dengan pendekatan Clean Architecture menggunakan Golang, Gin, GORM, PostgreSQL, dan Redis.

---

## 🚀 Tech Stack

- **Golang** (RESTful API)
- **Gin** (Web framework)
- **GORM** (ORM)
- **PostgreSQL** (Database)
- **Redis** (Cache + Queue)
- **JWT** (Authentication)
- Clean Architecture + Modular Structure

---

## 📦 Fitur

### 🔐 Authentication & Authorization
- JWT-based login
- Register & login endpoint

### 📝 Content Management
- Get All Posts (cached via Redis)
- Create, Update & Delete Post (by owner)
- Belongs to User & Category

### 🗂 Category Management
- Admin only access
- Create, Get, Update category

### 📈 Additional Features
- Redis Caching
- Transactional DB access
- Logging, Error Handling

---

## 📁 Folder Structure (Clean Architecture)

## 🧪 Postman Collection

📥 Import file `Blog Management API.postman_collection.json` untuk testing langsung di Postman.

---

## ⚙️ How to Run Locally

# 1. Clone Project
git clone https://github.com/your-username/blog-management-system.git
cd blog-management-system


# 2. Buat file .env
cat <<EOF > .env
PORT=9090

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=blog_management

REDIS_ADDR=localhost:6378
REDIS_PASS=
JWT_SECRET=secret
EOF

# 4. Install dependency
go mod tidy

# 5. Jalankan aplikasi
go run main.go


