# User Balance Management System

This system provides user authentication, balance management, and transaction history tracking using Laravel for the frontend and Golang for the backend.

## Features
- User Registration & Login (JWT Authentication)
- Check User Balance
- Top-up Balance (Simulated Payment Gateway)
- Perform Transactions (Simulated Payment Gateway)
- View Transaction History (Top-ups & Transfers)

## Prerequisites
Ensure you have the following installed:
- Docker & Docker Compose
- Golang (1.21)
- MySQL Server
- PHP (8.2+) (Laravel Framework)

## Installation

### Backend (Golang API)
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo.git
   cd backend
   ```

2. Create an `.env` file and configure database and JWT settings:
   ```sh
   SECRET_KEY=your_secret_key
   EXPIRATION_TIME=24 # Token expiration in hours
   DB_HOST=db
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=root
   DB_NAME=app_db
   ```

3. Start the backend server:
   ```sh
   go run main.go
   ```

### Frontend (Laravel)
1. Navigate to the frontend directory:
   ```sh
   cd frontend
   ```

2. Install dependencies:
   ```sh
   composer install
   npm install
   ```

3. Set up the `.env` file for Laravel:
   ```sh
   APP_URL=http://localhost
   API_URL=http://localhost:8080
   DB_CONNECTION=mysql
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_DATABASE=app_db
   DB_USERNAME=root
   DB_PASSWORD=root
   ```

4. Run database migrations:
   ```sh
   php artisan migrate
   ```

## Usage
### Register User
```sh
curl -X POST http://localhost:8080/api/auth/register -d '{"email": "user@example.com", "password": "password123"}' -H "Content-Type: application/json"
```

### Login User
```sh
curl -X POST http://localhost:8080/api/auth/login -d '{"email": "user@example.com", "password": "password123"}' -H "Content-Type: application/json"
```
Response:
```json
{
  "token": "your_jwt_token"
}
```

## Running with Docker
To run the application with Docker Compose:
```sh
docker-compose up -d
```

This will start the backend, frontend, and MySQL database.

## License
This project is licensed under the MIT License.

