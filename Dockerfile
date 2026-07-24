# ---------- Build Stage ----------
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

# ---------- Runtime Stage ----------
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/data ./data

RUN mkdir -p /app/data

CMD ["./app"]