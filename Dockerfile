FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:3.14
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
ENTRYPOINT ["./main"]
