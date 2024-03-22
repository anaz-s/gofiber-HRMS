---

# Sample Project with MongoDB and Fiber

This is a sample project showcasing how to build a RESTful API using Fiber (a web framework for Go) and MongoDB.

## Prerequisites

Before running this project, ensure you have the following installed on your machine:

- Go 1.16 or higher
- MongoDB
- Fiber (`github.com/gofiber/fiber/v3`)
- MongoDB Go Driver (`go.mongodb.org/mongo-driver`)

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/sample-project.git
```

2. Navigate to the project directory:

```bash
cd sample-project
```

3. Install dependencies:

```bash
go mod tidy
```

## Configuration

Make sure MongoDB is running on `localhost:27017`. If not, update the `uri` constant in `main.go` accordingly.

## Usage

To run the project, execute the following command:

```bash
go run main.go
```

By default, the server will listen on port `5000`. You can specify a different port using the `-port` flag:

```bash
go run main.go -port=8080
```

## API Endpoints

### Get All Users

```
GET /api/v1/users
```

Retrieves a list of all users.

### Add User

```
POST /api/v1/users
```

Adds a new user. The request body should be a JSON object with `firstName` and `lastName` fields.

### Update User

```
PUT /api/v1/users/:id
```

Updates an existing user by ID. The request body should contain the updated user object in JSON format.

### Delete User

```
DELETE /api/v1/users/:id
```

Deletes a user by ID.

## Project Structure

- `main.go`: Entry point of the application.
- `resources/user.go`: Defines the User struct and MongoDB instance.
  
## Contributing

Contributions are welcome! Please feel free to open a pull request or submit an issue if you find any bugs or want to suggest improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
