# SSL Checker Frontend (React)

Web interface built with **React**, **TypeScript**, and **Vite** to visualize SSL/TLS security results.

## 🚀 Development Setup

1.  **Install Dependencies:**
    ```bash
    npm install
    ```

2.  **Start Development Server:**
    ```bash
    npm run dev
    ```
    The app will be available at [http://localhost:5173](http://localhost:5173).

## 🔌 API Integration

The frontend communicates with the backend via the `/api/scan` endpoint.

* **Initiate New Scan:** `GET /api/scan?domain=example.com&new=true`
* **Check Status:** `GET /api/scan?domain=example.com`

The application handles the asynchronous nature of the SSL Labs API by implementing a polling mechanism.