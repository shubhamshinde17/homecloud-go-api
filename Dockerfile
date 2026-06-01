# Build

FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o homecloud-api \
    ./cmd/server

# Runtime

FROM scratch

WORKDIR /app

COPY --from=builder /app/homecloud-api .

EXPOSE 8080

CMD ["./homecloud-api"]
