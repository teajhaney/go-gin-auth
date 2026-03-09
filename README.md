# go-auth

A minimal Go authentication example using:

- Gin HTTP server
- MongoDB (official driver)
- JWT authentication (github.com/golang-jwt/jwt/v5)
- bcrypt password hashing

## 🧩 Project Structure

- `cmd/api` - application entrypoint and HTTP server setup
- `internal/app` - application dependencies and configuration
- `internal/server` - router and HTTP handlers
- `internal/middleware` - auth middleware (JWT verification)
- `internal/auth` - JWT creation/parsing and claim types
- `internal/user` - user model, repo, and service (registration/login)

## 🚀 Getting Started

### Prerequisites

- Go 1.25+
- MongoDB (local or Atlas)

### Setup

1. Copy `.env` and adjust values (especially the Mongo URI + JWT secret):

```bash
cp .env.example .env
# or edit .env directly
```

2. Install dependencies:

```bash
go mod download
```

3. Start the server:

```bash
go run ./cmd/api
```

The server will start on `http://localhost:3000` by default (from `.env`).

## 🧪 API Endpoints

### Health

- `GET /health`

Returns a simple `{ "ok": true }` response.

### Auth

#### Register

- `POST /register`
- Body (JSON):
  ```json
  { "email": "you@example.com", "password": "yourpass" }
  ```

#### Login

- `POST /login`
- Body (JSON):
  ```json
  { "email": "you@example.com", "password": "yourpass" }
  ```

### Protected Routes (require `Authorization: Bearer <token>`)

- `GET /api/files`
- `GET /api/products` (requires admin role)

## 🗝 Environment Variables

| Name | Description |
|------|-------------|
| `MONGODB_URI` | MongoDB connection string |
| `MONGO_DB_NAME` | Database name |
| `PORT` | HTTP server port |
| `JWT_SECRET` | Secret used to sign JWT tokens |
| `JWT_EXPIRATION` | Token expiration (e.g. `24h`) |

## ✅ Notes

- Passwords are hashed with `bcrypt` before saving.
- JWT tokens are signed with `HS256`.
- Role-based access is enforced via middleware.

---

If you want a specific feature added (refresh tokens, email verification, etc.), open an issue or update the docs accordingly.
