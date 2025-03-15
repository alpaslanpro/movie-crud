# Movie CRUD App

This is a simple Movie CRUD (Create, Read, Update, Delete) application built with **Go** using the following technologies:
- **UberFx** for Dependency Injection
- **Gin** for HTTP routing
- **GORM** for ORM with PostgreSQL (or MySQL)
- **JWT** for authentication and authorization

## Features
- User authentication with JWT tokens
- Basic CRUD operations for movies
- Graceful error handling with appropriate HTTP status codes
- Transaction handling with GORM
- Auto-generated API documentation (Swagger)
- Dockerized application for easy deployment

## Prerequisites
- Go 1.20+
- Docker & Docker Compose
- PostgreSQL (if running locally, otherwise Docker will handle it)

## Installation & Setup
1. **Clone the repository**:
```bash
git clone https://github.com/alpaslanpro/movie-crud.git
cd movie-crud
```

2. **Create a `.env` file** with your PostgreSQL connection details.

3. **Install dependencies**:
```bash
go mod tidy
```

4. **Run the application locally**:
```bash
go run main.go
```

5. **Run using Docker (Recommended)**:
```bash
docker-compose up --build
```

## API Endpoints
### Authentication
- `POST /login` - Generate JWT token for user.
- `POST /register` - Create a new user.

### Movies
- `POST /movies` - Create a new movie.
- `GET /movies` - Retrieve all movies.
- `GET /movies/:id` - Retrieve a specific movie by ID.
- `PUT /movies/:id` - Update an existing movie.
- `DELETE /movies/:id` - Delete a movie by ID.

## Future Improvements
- Optimize database queries using Gorm (if project gets bigger)
- Add logger to standardize output/error.
- Add unit and integration tests.
- Implement rate limiting and caching using Redis.
