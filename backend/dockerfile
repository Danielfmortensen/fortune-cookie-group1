# Build stage
FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o backend .

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/backend .

EXPOSE 9000

CMD ["./backend"]
