# Build stage
FROM golang:1.23.2-alpine AS builder

RUN apk update && apk add --no-cache ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Specify the path to the main.go file to build the application binary
RUN go build -o bin/application cmd/grpc_server/main.go

# Runner stage
FROM alpine AS runner

RUN apk update && apk add --no-cache ca-certificates libc6-compat bash && rm -rf /var/cache/apk/**

WORKDIR /opt

COPY --from=builder /opt/bin/application ./

# Set the entry point to the compiled application
CMD ["./application"]
