FROM golang:1.14-alpine AS builder
WORKDIR /app

COPY . .

RUN mkdir store

RUN go test -v --bench . --benchmem

RUN  GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /app/tobytes ./main.go

FROM scratch

COPY --from=builder /app/tobytes /app/tobytes

WORKDIR /app

ENTRYPOINT ["/app/tobytes"]
