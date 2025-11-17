# Project Layout

This document provides an overview of the project structure, files, and configuration.

## Project Structure

The project is organized to separate concerns, making it easier to navigate and maintain.

```
On progress    

```

### `/config`

This directory contains all application configuration logic.

*   **`config.go`**: Defines the Go structs that map to the configuration file. It uses the viper library to read settings from `config.yaml` and environment variables. It also provides a Google Wire provider set for dependency injection of the configuration.
*   **`config.yaml`**: The YAML file containing the actual configuration values for different environments.

## Configuration (`config.yaml`)

The `config.yaml` file stores all the configuration for the application. Viper allows these values to be overridden by environment variables. For example, `database.user` can be overridden by setting an environment variable `DATABASE_USER`.

### `app`

Contains general application settings.

*   `host`: The host address for the server (e.g., `"http://localhost"`).
*   `port`: The port on which the server will listen (e.g., `"8000"`).
*   `version`: The application version (e.g., `"1.0.0"`).

### `database`

Contains settings for the PostgreSQL database connection.

*   `host`: The database server host (e.g., `"127.0.0.1"`).
*   `port`: The database server port (e.g., `"5432"`).
*   `user`: The username for the database connection.
*   `passowrd`: The password for the database user. **Note:** There is a typo in this key. It should probably be `password`.
*   `name`: The name of the database.

### `redis`

Contains settings for the Redis connection.

*   `host`: The Redis server host (e.g., `"127.0.0.1"`).
*   `port`: The Redis server port (e.g., `"6379"`).
*   `password`: The password for Redis authentication. An empty string means no password is used.

### `jwt`

Contains settings for JSON Web Tokens (JWT).

*   `secret`: The secret key used to sign and verify JWTs.

### `signature`

*   `secret`: A secret for signing. **Note:** The `signature` section is present in `config.yaml`, but it is not defined in the `Config` struct in `config.go`. Therefore, its value is not being loaded into the application's configuration.