# Go Backend Boilerplate

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

this is a SIMPLE, yet a production-ready, feature-rich boilerplate for building robust and scalable backend services in Go. This project is built upon a foundation of modern best practices, including Clean Architecture, Dependency Injection, and a carefully selected set of high-quality libraries.

## Features

*   **High-Performance Web Framework**: Built with Fiber, a Go web framework inspired by Express.js, designed for speed and low memory overhead.
*   **Clean Architecture**: Enforces separation of concerns by dividing the application into distinct layers: `Handler`, `Usecase`, and `Repository`.
*   **Compile-Time Dependency Injection**: Utilizes Google Wire for robust, compile-time dependency injection, eliminating runtime reflection and improving performance.
*   **Powerful ORM**: Integrates GORM for elegant and efficient database interactions with PostgreSQL.
*   **Structured Logging**: Implements Go's native `log/slog` for structured, context-aware logging.
*   **Configuration Management**: Uses Viper to handle configuration from files (`config.yaml`) and environment variables.
*   **Robust CLI**: Powered by Cobra, providing commands for serving the application, running database migrations, and seeding data.
*   **Authentication**: Includes JWT-based authentication middleware to protect routes.
*   **Request Validation**: Employs `go-playground/validator` for easy and declarative validation of incoming request data.
*   **Database Migrations & Seeding**: Comes with built-in commands to manage your database schema and populate it with initial data.

## Project Structure

The project follows a logical and maintainable structure:

```
.
├── config/                 # Configuration files and loading logic
├── internal/               # Core application code (not for external import)
│   ├── app/http/           # HTTP server specifics
│   │   ├── handler/        # HTTP handlers (controllers)
│   │   ├── middleware/     # Custom middleware (e.g., auth)
│   │   └── router/         # Route definitions
│   ├── domain/             # Core business models and interface definitions
│   ├── repository/         # Data access layer (database interactions)
│   ├── usecase/            # Business logic layer
│   └── di/                 # Dependency Injection setup (Wire)
├── pkg/                    # Reusable packages safe for external use
│   ├── auth/               # JWT generation and validation
│   ├── cmd/                # Cobra CLI commands (serve, migrate, seed)
│   ├── database/           # DB connection, migration, and seeding logic
│   ├── logger/             # Structured logger setup
│   └── ...                 # Other shared utilities
├── go.mod                  # Go module definitions
└── main.go                 # Main application entry point
```

## Prerequisites

Before you begin, ensure you have the following installed:

*   **Go**: Version 1.21 or later.
*   **PostgreSQL**: A running instance of the PostgreSQL database.
*   **Wire CLI**: The code generator for Google Wire. Install it with:
    ```sh
    go install github.com/google/wire/cmd/wire@latest
    ```

## Getting Started

Follow these steps to get your local development environment up and running.

### 1. Clone the Repository

```sh
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
```

### 2. Configure the Application

Copy the example configuration file and update it with your local settings (database credentials, JWT secret, etc.).

```sh
cp config/config.yaml.example config/config.yaml
```

Now, open `config/config.yaml` and edit the values to match your environment.

**Important**: All configuration values in `config.yaml` can be overridden by environment variables. For example, to override the database password, you can set `DATABASE_PASSWORD=your_secret_password`.

### 3. Install Dependencies

Download the required Go modules.

```sh
go mod tidy
```

### 4. Generate Dependency Injection Code

Run the Wire CLI to generate the `wire_gen.go` files, which contain the dependency injection logic.

```sh
wire ./...
```

You should run this command whenever you add a new provider or change dependencies in `internal/di/wire.go`.

### 5. Run Database Migrations

Apply the database migrations to set up your tables.

```sh
go run main.go migrate
```

### 6. Seed the Database (Optional)

If you want to populate your database with initial data, run the seeder.

```sh
go run main.go seed
```

### 7. Run the Server

Start the HTTP server.

```sh
go run main.go serve
```

The application will be running on the port specified in your `config.yaml` (default is `8000`).

## Available CLI Commands

This application uses Cobra to provide a simple command-line interface.

| Command             | Description                                                              |
| ------------------- | ------------------------------------------------------------------------ |
| `serve`             | Starts the HTTP server.                                                  |
| `migrate`           | Runs all pending database migrations.                                    |
| `fresh-migrate`     | Drops all tables and re-runs all migrations. **Use with caution!**       |
| `seed`              | Populates the database with initial data defined in the seeder.          |

## API Endpoints

Here are some of the key API endpoints available in the boilerplate:

*   `POST /api/v1/dashboard/signin`: Authenticate as an administrator.
*   `GET /api/v1/dashboard/users`: Get all users (Protected).
*   `POST /api/v1/dashboard/users`: Create a new user (Protected).

*Note: Protected routes require a valid JWT `Authorization: Bearer <token>` header.*

## Contributing

Contributions are welcome! Please **feel free to submit a Pull Request.**

## License

This project is licensed under the MIT License. See the LICENSE file for details.