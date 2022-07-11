# go-test-class-api

This is a test RestAPI microservice, done in Go.

To start the server, you need to run `docker compose up -d`. Docker will create two postgres containers, one for production and other for testing, and one app container, called classes-service-production. This means that starting docker compose, the API will be ready to handle requests.

To run the application in production mode, you need to run `go run main.go`.

To run the application in test mode, you need to run `go run main.go -run-mode test`.

At .env file there are all the enviroment parameters. For testing:

- Database port: 54322
- API port: 5051

For production:

- Database port: 54321
- API port: 5051

If there's any change to .env params, it's important to rebuild the classes-service-production image.

There's a swagger.yml at doc/api with all the API documentation.

## Testing
To run the tests, I recommend to use the next command:

`go test ./tests/... -p 1`

I recommend to use the -p flag to 1, because if we use the concurrent tests, there's a possibility to found some issues relateds to the concurrencies and the different configurations.

### MockDb
The structs and funcs included at database/mock are used for testing. There are 3 types of tests:

- Model testing.
- API testing.
- Postgres testing.

Model and API testing are using mocked database, postgres is using postgres test container, because in that way I can control every different step of the processes, as I can validate if an error is caused by the database or by the code. Also, I can start coding without a database engine.

## API info
I've created a GET /classes endpoint, in order to get all the classes createds with their id. As it wasn't at the requirements, this endpoint doesn't have any queryString params, like limit, offset, sort or any other filter.

## Database
I'm using postgres because I've used it in other projects, so I didn't need to search for the connection string. Since the use of the DB was optional, I didn't do any research on which database engine would be best for this project.
