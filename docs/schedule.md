## 🗓️ FinBuddy Development Schedule

This document tracks the daily progress and plans for the development of FinBuddy.

---

### ✅ Week 1: Project Setup & Core API (May 20–26)

**May 20**

* Setup initial project directory and repo
* Dockerize basic Go app
* Create basic README structure
* Setup `main.go` and internal app layout

**May 21**

* Create `User` and `Pocket` models
* Implement initial migrations
* Setup connection to PostgreSQL via Docker Compose
* Start working on `GET /users` endpoint

**May 22**

* Implement `POST /users` and `POST /pockets`
* Finalize initial DB schema
* Improve error handling and response structuring

**May 23**

* Refactor service layers
* Add validator support to request payloads
* Improve modular structure of routers

**May 24–26** (Planned)

* Begin working on `Expense` and `Plan` models
* Add endpoint for adding daily expenses
* Add example data and seed script (optional)
* Write tests for user and pocket services

---

### 🧱 Week 2: Planning Logic & Dashboard API (May 27–June 2)

* Build logic to split monthly income into pockets
* Implement saving projection service
* Add endpoints:

  * `GET /dashboard`
  * `POST /plan`
* Start building Swagger or Postman documentation
* Polish and validate request/response formats

---

### 📦 Week 3: Summary Features & Yearly Targets (June 3–9)

* Add yearly expense summary service
* Add endpoints:

  * `GET /summary/yearly`
  * `POST /goal`
* Implement goal logic for future targets (e.g., house, vacation)
* Update README and architecture diagram

---

### 🔄 Ongoing / Future

* Add JWT auth & login endpoints
* Add frontend (optional, React/Tailwind)
* Unit and integration testing
* CI/CD setup
* Export/report feature (PDF/CSV)

---

*Last updated: May 20, 2025*
