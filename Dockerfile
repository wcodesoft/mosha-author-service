FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app/mosha-author-service

COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

ENV HTTP_PORT 8180
ENV GRPC_PORT 8181
ENV MONGO_DB_HOST "mongodb://localhost:27017"
ENV QUOTE_SERVICE_ADDRESS "localhost:8281"
ENV SENTRY_DSN ""
ENV SENTRY_SAMPLE_RATE "1.0"
ENV RELEASE_VERSION "dev"

WORKDIR /bin
COPY --from=builder /app/mosha-author-service/app .

CMD ["./app"]
EXPOSE 8180 8181