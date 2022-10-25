FROM golang:1.18-alpine3.16 as builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags '-w -s -extldflags "-static"' -o /app ./main

FROM scratch
COPY --from=builder /app /opt/backend
COPY config.toml /opt/config.toml
EXPOSE 8081/tcp
ENTRYPOINT ["/opt/backend"]
