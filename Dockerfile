FROM scratch

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY woodpecker-ntfy /

ENTRYPOINT ["/woodpecker-ntfy"]
