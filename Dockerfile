FROM golang:1.23-alpine AS builder

WORKDIR /
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/confirm ./cmd/
RUN ls -l /app

FROM alpine:latest 

WORKDIR /app
COPY --from=builder /app/confirm .
COPY --from=builder /migrations ./migrations

ENTRYPOINT ["./confirm"]