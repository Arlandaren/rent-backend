FROM golang:1.23.2-alpine AS builder

RUN apk update && apk add --no-cache ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o bin/grpc_server cmd/grpc_server/main.go

RUN go build -o bin/migrator cmd/migrator/main.go

FROM alpine AS runner

RUN apk update && apk add --no-cache ca-certificates libc6-compat bash && rm -rf /var/cache/apk/**

WORKDIR /opt

COPY --from=builder /opt/migrations ./migrations
COPY --from=builder /opt/bin/grpc_server ./
COPY --from=builder /opt/bin/migrator ./

CMD ["./grpc_server"]
