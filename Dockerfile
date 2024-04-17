FROM golang:alpine AS builder
RUN CGO_ENABLED=0 go build -o server

EXPOSE 9000

CMD ["server"]
