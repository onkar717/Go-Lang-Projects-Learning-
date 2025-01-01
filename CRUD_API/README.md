# Movie Management API

This is a simple RESTful API built in Go (Golang) to manage a collection of movies. The API supports CRUD (Create, Read, Update, Delete) operations for movies, utilizing the `gorilla/mux` package for routing.

## Features

- Retrieve all movies
- Retrieve a single movie by ID
- Add a new movie
- Update an existing movie
- Delete a movie

## Prerequisites

- Go (version 1.16 or higher)
- `gorilla/mux` package

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/movies-api.git
   cd movies-api
   ```

2. Install dependencies:

   ```bash
   go get -u github.com/gorilla/mux
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. The server will start on `http://localhost:8080`.

## API Endpoints

### Get All Movies
- **URL**: `/movies`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "id": "1",
      "isbn": "438227",
      "title": "Movie One",
      "director": {
        "firstname": "Onkar",
        "lastname": "Shelke"
      }
    }
  ]
  ```

### Get Movie by ID
- **URL**: `/movies/{id}`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "id": "1",
    "isbn": "438227",
    "title": "Movie One",
    "director": {
      "firstname": "Onkar",
      "lastname": "Shelke"
    }
  }
  ```

### Add a New Movie
- **URL**: `/movies`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "isbn": "123456",
    "title": "New Movie",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  }
  ```
- **Response**:
  ```json
  {
    "id": "random-generated-id",
    "isbn": "123456",
    "title": "New Movie",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  }
  ```

### Update a Movie
- **URL**: `/movies/{id}`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "isbn": "123456",
    "title": "Updated Movie",
    "director": {
      "firstname": "Jane",
      "lastname": "Doe"
    }
  }
  ```
- **Response**:
  ```json
  {
    "id": "1",
    "isbn": "123456",
    "title": "Updated Movie",
    "director": {
      "firstname": "Jane",
      "lastname": "Doe"
    }
  }
  ```

### Delete a Movie
- **URL**: `/movies/{id}`
- **Method**: `DELETE`
- **Response**:
  ```json
  [
    {
      "id": "2",
      "isbn": "438228",
      "title": "Movie Two",
      "director": {
        "firstname": "Rohit",
        "lastname": "Shelke"
      }
    }
  ]
  ```

## Project Structure

```
movies-api/
├── main.go
└── README.md
```

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
