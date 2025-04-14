
# BeamBridge 🚀

[![Go](https://img.shields.io/badge/Go-1.24-00ADD8.svg?logo=go)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Compose-blue.svg?logo=docker)](https://www.docker.com)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

**BeamBridge** is a modern file transfer system built in Go, designed to send photos, videos, and iPhone Live Photos from your phone to your PC. It supports **chunked uploads**, real-time progress tracking with **SSE**, and follows a clean, maintainable **Hybrid Hexagonal-Clean Architecture** – ideal for learning and applying backend best practices! 📸💾

---

## ✨ Features

- 🚚 **Chunked Uploads** – ideal for large files.
- ⚡ **Non-Chunked Uploads** – simple and fast for small files.
- 📷 **iPhone Live Photos Support** – send image + video together.
- 📡 **Real-Time Progress** – via Server-Sent Events (SSE).
- 🗃️ **Metadata Storage** – PostgreSQL with JSONB fields.
- ⚙️ **Redis Cache** – handles sessions and upload progress.
- ♻️ **Hot Reload with Air** – fast development.
- 🐳 **Docker Ready** – run everything with `docker-compose up`.

---

## 🛠️ Tech Stack

| Technology     | Purpose                      |
|----------------|-------------------------------|
| Go 1.24        | Core backend language         |
| PostgreSQL     | Database for metadata         |
| Redis          | Caching sessions/progress     |
| Docker         | Containerization              |
| Air            | Hot reload for development    |
| UUID           | Unique upload identifiers     |

---

## 🚀 Getting Started

### 📦 Requirements

- [Docker + Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/)
- [Go 1.24+](https://go.dev/doc/install) (optional – if running without Docker)

---

### 🧭 Steps

```bash
# Clone the project
git clone https://github.com/BrunoGuimaraesSilva/beambridge.git
cd beambridge

# Build and run containers
docker-compose up --build
```

- Go service runs at `http://localhost:8080`
- PostgreSQL at `5432`, Redis at `6379`
- DB initialized with `init.sql` (tables: `files`, `live_photos`)
- Hot reload with `air`

---

## 🔬 API Testing

### 📤 Full Upload
```bash
echo "test" > test.jpg
curl -X POST -F "fileName=test.jpg" -F "uploadId=123" -F "file=@test.jpg" http://localhost:8080/upload/file
```

### 📤 Chunked Upload
```bash
curl -X POST -F "fileName=test.jpg" -F "uploadId=123" -F "chunk=@test.jpg" http://localhost:8080/upload/chunk
```

### 📶 Real-Time Progress via SSE
```bash
curl http://localhost:8080/progress?uploadId=123
```

---

## 🧬 Project Structure

```
beambridge/
├── cmd/                 # main.go
├── internal/
│   ├── domain/          # Entities: File, LivePhoto
│   ├── usecase/         # Business logic
│   ├── port/            # Interfaces
│   ├── adapter/         # Redis, PostgreSQL, Storage, HTTP
│   ├── config/          # Environment variables
│   └── logger/          # Centralized logging
├── init.sql             # Database schema
├── Dockerfile           # Go app container
├── docker-compose.yml   # PostgreSQL, Redis, App
├── .air.toml            # Hot reload config
├── .gitignore
├── go.mod / go.sum
└── README.md            # You're here ✨
```

---

## ⚙️ How It Works

> BeamBridge separates concerns using clean principles. Upload sessions are tracked using Redis, files are stored and indexed with PostgreSQL, and progress is streamed live via SSE. Chunked uploads allow reliable transfers even on unstable networks.