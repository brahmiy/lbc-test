FROM golang:1.19 as builder

COPY . /go/src/

RUN cd /go/src \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -ldflags="-s -w" -o /app .

FROM gcr.io/distroless/base

COPY --from=builder /app .

CMD ["/app"]
