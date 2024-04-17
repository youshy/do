FROM golang:alpine AS builder
COPY . .
RUN CGO_ENABLED=0 go build -o server

EXPOSE 9000

CMD ["server"]
