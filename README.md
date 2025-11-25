# KageVault API

KageVault API is a high-performance, secure password generation service built with Go and the Fiber web framework. It provides a simple and efficient way to generate strong, customizable passwords for your applications.

##  Features

- **Customizable Password Generation**: Control length, character sets (digits, symbols, uppercase).
- **High Performance**: Built on top of [Fiber](https://gofiber.io/), one of the fastest Go web frameworks.
- **Docker Ready**: Includes Dockerfile and docker-compose setup for easy deployment.
- **Health Checks**: Built-in health check endpoint for monitoring.

## Tech Stack

- **Language**: Go 1.25+
- **Framework**: Fiber v2
- **Configuration**: Godotenv

##  Installation & Running

### Prerequisites

- Go 1.25 or higher
- Docker (optional, for containerized deployment)

### Running Locally

1. **Clone the repository**
   ```bash
   git clone https://github.com/pavelc4/kage-vault-go.git
   cd kage-vault-go
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run cmd/api/main.go
   ```
   The server will start on port `3000` (default).

### Running with Docker

1. **Build and start the container**
   ```bash
   docker-compose up -d --build
   ```
   The API will be available at `http://localhost:3000`.

##  API Endpoints

### 1. Service Info
Returns basic information about the service.

- **URL**: `/`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "service": "KageVault API",
      "version": "1.0.0"
    }
  }
  ```

### 2. Health Check
Check the health and uptime of the service.

- **URL**: `/api/health`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "status": "healthy",
      "uptime": "1m30s",
      "service": "KageVault API"
    }
  }
  ```

### 3. Generate Password
Generate a secure password with custom options.

- **URL**: `/api/password`
- **Method**: `GET`
- **Query Parameters**:
  - `length` (int): Length of the password (default: `12`)
  - `digits` (bool): Include numbers (default: `true`)
  - `symbols` (bool): Include special characters (default: `true`)
  - `uppercase` (bool): Include uppercase letters (default: `true`)

- **Example Request**:
  ```
  GET /api/password?length=16&symbols=false
  ```

- **Response**:
  ```json
  {
    "success": true,
    "message": "Password generated successfully",
    "data": {
      "password": "GeneratedPassword123",
      "length": 16
    }
  }
  ```

##  Configuration

The application can be configured using environment variables or a `.env` file.

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT`   | Server port | `3000`  |

##  License

[MIT](LICENSE)
