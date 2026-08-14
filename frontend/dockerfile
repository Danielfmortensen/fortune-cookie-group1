# Build stage
FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o frontend .

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/frontend .
COPY --from=build /app/static ./static
COPY --from=build /app/templates ./templates

EXPOSE 8080

CMD ["./frontend"]
