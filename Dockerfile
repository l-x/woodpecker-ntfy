FROM alpine
COPY woodpecker-ntfy /
ENTRYPOINT ["/woodpecker-ntfy"]
