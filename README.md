# FinBuddy 🧾

FinBuddy is a personal financial planning app built with Go and PostgreSQL. It helps individuals and families plan monthly expenses, track daily spending, and allocate savings toward long-term goals.

## 🚀 Features (Planned)

- ✅ Dashboard overview
- ✅ Monthly planning with customizable categories (needs, wants, savings)
- ✅ Daily expense tracking
- ✅ Saving projection for a year
- ✅ Yearly expense summary
- ✅ "Pocket" system for monthly budget allocations
- ✅ Saving splits into investment options
- ✅ Monthly, yearly, lifetime saving targets (e.g. house, travel, Hajj)

## 🏗 Tech Stack

- **Backend:** Go
- **Database:** PostgreSQL (via Docker)
- **Dev Environment:** WSL2 (Windows Subsystem for Linux)
- **ORM / SQL:** TBD
- **API:** RESTful (JSON)

## 🛠 Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/finbuddy.git
cd finbuddy
```

### 2. Setup Environment
Create a `.env` file:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=finbuddy_user
DB_PASSWORD=finbuddy_pass
DB_NAME=finbuddy_db
```

### 3. Run PostgreSQL with Docker
```bash
docker run --name finbuddy-db \
  -e POSTGRES_USER=finbuddy_user \
  -e POSTGRES_PASSWORD=finbuddy_pass \
  -e POSTGRES_DB=finbuddy_db \
  -p 5432:5432 \
  -d postgres
```

### 4. Run the App
```bash
go run main.go
```

Visit: [http://localhost:8080/healthz](http://localhost:8080/healthz)

### 5. Run the Test
```bash
docker build --target test .
```

## 📁 Project Structure

```
finbuddy/
├── main.go
├── .env
├── go.mod
├── go.sum
├── README.md
└── ...
```

## 💡 Development Schedule

See [`docs/schedule.md`](docs/schedule.md) _(coming soon)_ for daily breakdowns and progress logs.

---

## 🧑‍💻 Author

Built by [@idilhaq](https://github.com/idilhaq) to improve Go backend and product development skills.

---

## 📜 License

MIT License
