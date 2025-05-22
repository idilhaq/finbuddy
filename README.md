# FinBuddy 📟

FinBuddy is a personal financial planning app built with Go and PostgreSQL. It helps individuals and families plan monthly expenses, track daily spending, and allocate savings toward long-term goals.

## 🚀 Features (Planned)

* ✅ Dashboard overview
* ✅ Monthly planning with customizable categories (needs, wants, savings)
* ✅ Daily expense tracking
* ✅ Saving projection for a year
* ✅ Yearly expense summary
* ✅ "Pocket" system for monthly budget allocations
* ✅ Saving splits into investment options
* ✅ Monthly, yearly, lifetime saving targets (e.g. house, travel, Hajj)

## 🏗 Tech Stack

* **Backend:** Go (Gin framework)
* **Database:** PostgreSQL (via Docker)
* **Dev Environment:** WSL2 (Windows Subsystem for Linux) + Docker + `air` for hot reload
* **ORM / SQL:** GORM (with auto-migrations)
* **API:** RESTful (JSON) with planned OpenAPI/Swagger documentation
* **Build:** Multi-stage Dockerfile with separate `dev` and `prod` stages
* **Dev Tools:** `make`, `air`, Docker Compose, `.air.toml`

## 🛠 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/finbuddy.git
cd finbuddy
```

---

### 2. Create Environment Variables

Create a `.env` file in the root project directory:

```
DB_HOST=db
DB_PORT=5432
DB_USER=finbuddy_user
DB_PASSWORD=finbuddy_pass
DB_NAME=finbuddy_db
```

> ✅ These are used by the API container to connect to the DB container.

---

### 3. Start Development Environment (Hot Reload with Air)

```bash
make dev-up
```

> This uses `docker-compose.override.yml` and mounts your source code with `air` for hot reload.

Then open: [http://localhost:8080/healthz](http://localhost:8080/healthz)

---

### 4. Run Tests

```bash
make test
```

Or via Docker:

```bash
docker build --target test .
```

---

### 5. Stop Dev Environment

```bash
make dev-down
```

---

## 📁 Project Structure

```
finbuddy/
├── cmd/
│   └── api/
│       └── main.go         # API entry point
├── internal/
│   ├── db/                 # DB setup and models
│   ├── router/             # Gin router setup
│   └── handler/            # HTTP handlers
├── docker-compose.yml
├── docker-compose.override.yml
├── Dockerfile
├── .env
├── .air.toml
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 💡 Development Schedule

See [`docs/schedule.md`](docs/schedule.md) *(coming soon)* for daily breakdowns and progress logs.

---

## 🧑‍💻 Author

Built by [@idilhaq](https://github.com/idilhaq) to improve Go backend and product development skills.

---

## 📜 License

MIT License
