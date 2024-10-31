# Chirpy

## Table of Contents

- [About](#about)
- [API](#api)
  - [Users](#users)
  - [Auth](#auth)
  - [Chirps](#chirps)
  - [Admin](#admin)
  - [General](#general)
  - [Polka](#polka)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#clone-the-repo)
  - [Database Setup](#database-setup)

## About

Chirpy is a social network similar to Twitter. This project focused on Chirpy's backend, including:

- Building an HTTP server in Go
- Creating a RESTful API for clients to interact with, supporting CRUD operations
- Implementing authentication/authorization with cryptography libraries
- Working with JWTs, hashing, API keys
- Using SQL to store and retrieve data from a Postgres database

## API

### Users

User resource:

```json
{
  "id": "37393de8-83f9-464d-bf88-3c9768057e1e",
  "created_at": "2024-10-30T16:26:43.530893Z",
  "updated_at": "2024-10-30T16:26:43.530893Z",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

#### `POST /api/users`

Creates a new user.

Request body:

```json
{
  "email": "user@example.com",
  "password": "123456"
}
```

Response body:

```json
{
  "id": "37393de8-83f9-464d-bf88-3c9768057e1e",
  "created_at": "2024-10-30T16:26:43.530893Z",
  "updated_at": "2024-10-30T16:26:43.530893Z",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

#### `PUT /api/users`

Updates a user's email and password. Access token required.

Headers:

```json
{
  "Authorization": "Bearer <token>"
}
```

Request body:

```json
{
  "email": "new_email@example.com",
  "password": "new_password"
}
```

Response body:

```json
{
  "id": "37393de8-83f9-464d-bf88-3c9768057e1e",
  "created_at": "2024-10-30T16:26:43.530893Z",
  "updated_at": "2024-10-30T17:02:15.720771Z",
  "email": "new_email@example.com",
  "is_chirpy_red": false
}
```

### Auth

#### `POST /api/login`

Validates a user's password and creates an access and refresh token.

Request body:

```json
{
  "email": "user@example.com",
  "password": "123456"
}
```

Response body:

```json
{
  "id": "37393de8-83f9-464d-bf88-3c9768057e1e",
  "created_at": "2024-10-30T16:26:43.530893Z",
  "updated_at": "2024-10-30T16:26:43.530893Z",
  "email": "user@example.com",
  "is_chirpy_red": false,
  "token": "<token>",
  "refresh_token": "<refresh_token>"
}
```

#### `POST /api/refresh`

Generates a new access token. Refresh token required.

Headers:

```json
{
  "Authorization": "Bearer <refresh_token>"
}
```

Response body:

```json
{
  "token": "<token>"
}
```

#### `POST /api/revoke`

Revokes a refresh token. Refresh token required.

Responds with 204 status code.

Headers:

```json
{
  "Authorization": "Bearer <refresh_token>"
}
```

### Chirps

Chirp resource:

```json
{
  "id": "1d7c2a5d-ecc3-4da5-9311-56133ac8939a",
  "created_at": "2024-10-30T18:29:12.896088Z",
  "updated_at": "2024-10-30T18:29:12.896088Z",
  "body": "this is a chirp!",
  "user_id": "d33b88e9-19fb-426c-80a4-e2667c8180d3"
}
```

#### `POST /api/chirps`

Creates a new chirp and writes it to the database.

Chirp length must not exceed 140 characters or contain bad words (kerfuffle, sharbert, fornax).

Access token required.

Headers:

```json
{
  "Authorization": "Bearer <token>"
}
```

Request body:

```json
{
  "body": "This is a chirp!"
}
```

Response body:

```json
{
  "id": "1d7c2a5d-ecc3-4da5-9311-56133ac8939a",
  "created_at": "2024-10-30T18:29:12.896088Z",
  "updated_at": "2024-10-30T18:29:12.896088Z",
  "body": "This is a chirp!",
  "user_id": "d33b88e9-19fb-426c-80a4-e2667c8180d3"
}
```

#### `GET /api/chirps`

Returns an array of chirps with optional query parameters (sort, author_id).

Sort values can be "asc" or "desc".

Examples:

```
GET http://localhost:8080/api/chirps
GET http://localhost:8080/api/chirps?sort=asc
GET http://localhost:8080/api/chirps?sort=desc
GET http://localhost:8080/api/chirps?sort=asc&author_id=2
```

#### `GET /api/chirps/{chirpID}`

Returns a single chirp by its ID.

#### `DELETE /api/chirps/{chirpID}`

Deletes a single chirp by its ID. Access token required for author validation.

Responds with 204 status code.

Headers:

```json
{
  "Authorization": "Bearer <token>"
}
```

### Admin

#### `GET /admin/metrics`

Returns a welcome page with the count of visits to the app.

#### `POST /admin/reset`

Resets the count of app visits and deletes all users from the database.

Responds with 200 status code.

### General

#### `GET /app`

Serves Chirpy's home page.

#### `GET /app/assets/logo.png`

Serves Chirpy's logo.

#### `GET /api/healthz`

Checks server status.

Responds with 200 status code.

### Polka

#### `POST /api/polka/webhooks`

Upgrades the user to Chirpy Red if the event is "user.upgraded". Polka API key required.

Responds with 204 status code.

Headers:

```json
{
  "Authorization": "ApiKey <POLKA_KEY>"
}
```

Request body:

```json
{
  "data": {
    "user_id": "5d006c11-dfc7-41f9-9c68-fe0d716f26b3"
  },
  "event": "user.upgraded"
}
```

## Getting Started

### Prerequisites

- Go - [official download page](https://go.dev/doc/install)
- Postgres
  - Mac OS with brew:
    ```bash
    brew install postgresql@16
    ```
  - Linux/WSL:
    ```bash
    sudo apt update
    sudo apt install postgresql postgresql-contrib
    ```
- Goose
  ```bash
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```

### Clone the repo

1. Clone the repository to your local machine:

```bash
git clone https://github.com/mshortcodes/chirpy_new
```

2. Change into the project directory:

```bash
cd chirpy_new
```

### Database setup

1. Start the Postgres server.

Mac:

```bash
brew services start postgresql
```

Linux:

```bash
sudo service postgresql start
```

2. Create a new database. You may be prompted to enter a Postgres password if it hasn't been setup.

Mac:

```bash
psql postgres
```

Linux:

```bash
sudo -u postgres psql
```

In psql:

```sql
CREATE DATABASE chirpy;
```

3. Create a .env file. Chirpy requires setting up a couple environment variables.

```
touch .env
```

The .env file should have this structure:

```
DB_URL="<connection_string>"
PLATFORM="dev"
JWT_SECRET="<signing_key>"
POLKA_KEY="<api_key>"
```

Replace <connection_string> with your database connection string:

protocol://username:password@host:port/database?sslmode=disable

Replace <signing_key> with a randomly generated string:

```bash
openssl rand -base64 64
```

Replace <api_key> with another generated string:

```bash
openssl rand -hex 32
```

4. Run database migrations.

```bash
cd sql/schema
goose postgres <connection_string> up
```

5. Run the server.

```bash
cd ../..
go run .
```
