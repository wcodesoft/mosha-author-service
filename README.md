# mosha-author-service

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/e648d9edadb04e779a3d17325c437813)](https://app.codacy.com/gh/wcodesoft/mosha-author-service/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

Author microservice used in Mosha.

## Database

The main database used in the service is MongoDB. It's used to store the authors. To deploy it locally, run:

```bash
docker run --name mongo -p 27017:27017 -d mongodb/mongodb-community-server:latest 
```

## Docker

To build the container image, run:

```bash
docker build -t mosha-author-service .
```

After that to run the container, run:

```bash
docker run --name mosha-author-service -e MONGO_DB_HOST="mongodb://localhost:27017" --net=bridge -p 8180:8180 -d mosha-author-service
```

## gRPC

The communication between services is done using gRPC. To regenerate the gRPC code, run:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/author.proto
```

## Tests

Unit tests are written using [goconvey](https://smartystreets.github.io/goconvey/) library in go for more fluent test development. All
fake data in tests is generated using [gofakeit](https://github.com/brianvoe/gofakeit/) library.

To run all tests execute the following command in the root folder:

```bash
go test -v ./...
```