FROM golang:1.20-alpine as builder
WORKDIR /go/src/github.com/chappleg/request-printer
COPY request-printer.go go.mod ./
RUN go build -ldflags "-s -w"

FROM alpine
COPY --from=builder /go/src/github.com/chappleg/request-printer/request-printer /bin/
EXPOSE 80
CMD ["request-printer"]