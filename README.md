## URL Shortener

URL shortener application written in Go using the [hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) pattern. It uses MySQL as its database.

### Prerequisites

- [Go](https://golang.org/) (version 1.16 or later)
- [MySQL](https://www.mysql.com/)

### Database Credentials

Before running the application, you need to set up the database connection. Update the database configuration in the `start.sh` file.

### Running the Application

To run the application, execute the following command: `chmod +x start.sh && ./start.sh`

### Testing

To run the tests, execute the following command: `go test ./...`

### Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### License

This project is licensed under the [MIT License](./LICENSE).

### Copyright

(c) 2023 [Finbarrs Oketunji](https://finbarrs.eu). All Rights Reserved.