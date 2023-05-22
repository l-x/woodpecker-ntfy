FROM golang:1.20.4-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY plugin/*.go ./plugin/

RUN CGO_ENABLED=0 GOOS=linux go build -o /woodpecker-ntfy

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder woodpecker-ntfy /

ENTRYPOINT ["/woodpecker-ntfy"]
