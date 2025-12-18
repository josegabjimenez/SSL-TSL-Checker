# SSL/TLS Security Checker

This is an application that analyzes the SSL/TLS security of any given domain using the [Qualys SSL Labs API v3](https://github.com/ssllabs/ssllabs-scan/blob/master/ssllabs-api-docs-v3.md).

## 🚀 Quick Start (Docker)

The easiest way to run the application is with Docker Compose. This starts both the Go backend and the React frontend.

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd ssl-tsl-checker
    ```

2.  **Start the services:**
    ```bash
    docker-compose up --build
    ```

3.  **Access the application:**
    * **Frontend:** [http://localhost:5173](http://localhost:5173)
    * **Backend API:** [http://localhost:8080](http://localhost:8080)

## 🛠️ Architecture

* **Backend (Go):** A HTTP server that interfaces with SSL Labs.
* **Frontend (React + TypeScript):** A clean interface that polls the backend for scan progress and displays the final report.
* **Infrastructure:** Fully containerized with Docker, featuring hot-reloading for both backend (Air) and frontend (Vite) to facilitate development.