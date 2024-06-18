## Go Fiber JWT Authentication

This is a tutorial project that demonstrates how to implement JWT (JSON Web Token) authentication using the Go Fiber web framework.

### Features

- User registration
- User login
- Authenticated user retrieval
- User logout

### Technologies Used

- Go
- Go Fiber
- GORM (Object-Relational Mapping)
- MySQL
- JWT (JSON Web Tokens)
- Bcrypt (Password hashing)

### Project Structure

The project is structured as follows:

```
.
├── controllers
│   └── authController.go
├── database
│   └── connection.go
├── main.go
├── models
│   └── user.go
└── routes
    └── routes.go
```

- `main.go`: The entry point of the application, responsible for setting up the Fiber app, connecting to the database, and registering the routes.
- `database/connection.go`: Handles the connection to the MySQL database using GORM.
- `models/user.go`: Defines the User model and its properties.
- `routes/routes.go`: Defines the application routes and maps them to the corresponding controllers.
- `controllers/authController.go`: Implements the logic for user registration, login, user retrieval, and logout.

### Getting Started

1. Install Go and set up your development environment.
2. Install the required dependencies:
   - Go Fiber: `go get github.com/gofiber/fiber/v2`
   - GORM: `go get gorm.io/gorm` and `go get gorm.io/driver/mysql`
   - JWT: `go get github.com/golang-jwt/jwt/v5`
   - Bcrypt: `go get golang.org/x/crypto/bcrypt`
3. Create a MySQL database named `go_auth`.
4. Update the database connection details in `database/connection.go` to match your MySQL configuration.
5. Run the application:
   ```
   go run main.go
   ```
6. The application will start running on `http://localhost:8000`.

### API Endpoints

- `POST /api/register`: Register a new user.
- `POST /api/login`: Log in a user and receive a JWT cookie.
- `GET /api/user`: Retrieve the authenticated user's information.
- `POST /api/logout`: Log out the user by removing the JWT cookie.

### Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

### License

This project is licensed under the [MIT License](LICENSE).

Citations:
[1] https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/14001646/f0da3038-3688-4712-8d9a-f6ace3c46620/Go Fiber JWT Authentication - Tutorial - freeCodeC 1f6beaf4b9e842f8b89180ee1272f962.pdf