# Build stage
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=1 \
    GO111MODULE=on
RUN apk update && apk add --no-cache gcc musl-dev sqlite sqlite-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o loanapp .

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/loanapp .
COPY --from=builder /app/loan_app.db ./

EXPOSE 8080

CMD ["./loanapp"]
