# mosha-author-service

[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/e648d9edadb04e779a3d17325c437813)](https://app.codacy.com/gh/wcodesoft/mosha-author-service/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/e648d9edadb04e779a3d17325c437813)](https://app.codacy.com/gh/wcodesoft/mosha-author-service/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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

## Tests

Unit tests are written using https://smartystreets.github.io/goconvey/ library in go for more fluent test development.
All
fake data in tests is generated using https://github.com/brianvoe/gofakeit/ library.
