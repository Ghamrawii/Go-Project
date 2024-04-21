# Event Management API

This project is an API service for managing events. It allows users to create, update, delete, and retrieve event information through defined routes. The API is built using Go (Golang) and leverages the Gin framework for building the API endpoints, JWT for authentication, and SQLite for data storage.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (Golang) installed on your machine. You can download it from [here](https://golang.org/dl/).
- SQLite installed on your machine. You can download it from [here](https://www.sqlite.org/download.html).

### Installation

1. Clone the repository:

2. Navigate to the project directory:

3. Install the dependencies:
   ```
   go mod download
   ```

## Running the API

To start the API server, run the following command in the project directory:

```
go run .
```

The API will start running on `http://localhost:8080`.

## API Endpoints

### Events

- `GET /events`: Retrieve all events.
- `GET /events/:id`: Retrieve a specific event by ID.
- `POST /events`: Create a new event. Requires authentication.
- `PUT /events/:id`: Update a specific event by ID. Requires authentication.
- `DELETE /events/:id`: Delete a specific event by ID. Requires authentication.

### User Management

- `POST /user/signup`: Register a new user.
- `POST /user/login`: Authenticate a user and return a JWT token.

### Event Registration

- `POST /events/:id/register`: Register for a specific event. Requires authentication.
- `DELETE /events/:id/register`: Cancel registration for a specific event. Requires authentication.
