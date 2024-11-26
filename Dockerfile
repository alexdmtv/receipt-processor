FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ENV PORT=8080

EXPOSE ${PORT}
CMD ["./server"]
