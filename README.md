# mosha-author-service

[![DeepSource](https://app.deepsource.com/gh/wcodesoft/mosha-author-service.svg/?label=active+issues&token=nviNNhLutptceoSk2YUDaoDm)](https://app.deepsource.com/gh/wcodesoft/mosha-author-service/?ref=repository-badge)
[![DeepSource](https://app.deepsource.com/gh/wcodesoft/mosha-author-service.svg/?label=resolved+issues&token=nviNNhLutptceoSk2YUDaoDm)](https://app.deepsource.com/gh/wcodesoft/mosha-author-service/?ref=repository-badge)

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