# url-shortener-go

A high-performance URL shortening service built with **Go**, **Redis**, and **React (Vite)**.

---

## Getting Started

### Prerequisites
Ensure you have the following installed on your machine:
* [Docker](https://www.docker.com/get-started)
* [Docker Compose](https://docs.docker.com/compose/install/)

### 1. Project Structure
Ensure your environment files are located in the following directories:

**Backend Configuration (`backend/api/.env`):**
```env
DB_ADDR="db:6379"
DB_PASS=""
APP_PORT=":3000"
DOMAIN="localhost:3000"
API_QUOTA=10
FRONTEND_DOMAIN="http://localhost:5173"

```

**Frontend Configuration (`frontend/.env`):**

```env
VITE_GO_SHORTEN_URL=http://localhost:3000/api/shorten
VITE_GO_DOMAIN=http://localhost:3000

```

---

### 2. Launching the Project

To build the images and start all services (Frontend, Backend, and Redis) simultaneously, run:

```bash
docker-compose up --build

```

Once the containers are healthy:

* **Frontend UI:** [http://localhost:5173](https://www.google.com/search?q=http://localhost:5173)
* **Backend API:** [http://localhost:3000](https://www.google.com/search?q=http://localhost:3000)
* **Redis Store:** Internal port `6379` (Accessible via the `db` service name)

---

## Tech Stack & Architecture

* **Frontend:** React + Vite for a blazing fast, modern user interface.
* **Backend:** Go (Golang) handling API requests, validation, and redirection logic.
* **Database:** Redis for lightning-fast key-value storage and TTL (Time-To-Live) management.
* **Containerization:** Docker Compose for seamless multi-service orchestration.

---

## Useful Commands

| Action | Command |
| --- | --- |
| **Start Services** | `docker-compose up -d` |
| **Stop Services** | `docker-compose down` |
| **View Logs** | `docker-compose logs -f` |
| **Rebuild** | `docker-compose up --build` |

---

## API Limits

The current configuration is set to an **API_QUOTA** of `10` requests per period to prevent service abuse. You can modify this in the `backend/api/.env` file.
