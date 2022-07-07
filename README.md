# go-test-class-api

This is a test RestAPI microservice, done in Go.

To start the server, you need to run `docker compose up -d`. Docker will create two postgres containers, one for production and other for testing.

To run the application in production mode, you need to run `go run main.go`.

To run the application in test mode, you need to run `go run main.go -run-mode test`.

At .env file there are all the enviroment parameters. For testing:

- Database port: 54322
- API port: 5051

For production:

- Database port: 54321
- API port: 5051

There's a swagger.yml at doc/api with all the API documentation.

To run the tests, I recommend to use the next command:

`go test ./tests/... -p 1`

I recommend to use the -p flag to 1, because if we use the concurrent tests, there's a possibility to found some issues relateds to the concurrencies and the different configurations.
