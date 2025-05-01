# Go CRUD Application

A RESTful CRUD (Create, Read, Update, Delete) application built with Go, Gin framework, and PostgreSQL database.

## Features

- RESTful API endpoints for CRUD operations
- PostgreSQL database integration using GORM
- Environment variable configuration
- Clean architecture with separate routes, controllers, and models
- Automatic database migrations

## Prerequisites

- Go 1.24.2 or higher
- PostgreSQL database
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/LikhithMar14/go-crud-app.git
cd go-crud-app
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```env
DB_HOST=your_db_host
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=your_db_port
PORT=8080
```

4. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

## Project Structure

```
.
├── config/         # Database configuration
├── controllers/    # Request handlers
├── models/         # Data models
├── routes/         # API routes
├── main.go         # Application entry point
└── go.mod          # Go module file
```

## API Endpoints

The application provides the following RESTful endpoints:

- `GET /books` - Get all books
- `GET /books/:id` - Get a specific book
- `POST /books` - Create a new book
- `PUT /books/:id` - Update a book
- `DELETE /books/:id` - Delete a book

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader
- [PostgreSQL](https://www.postgresql.org/) - Database

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 