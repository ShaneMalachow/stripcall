FROM golang:1.12.1-alpine AS builder
COPY . /multistage
WORKDIR /multistage
ENV GO111MODULE=on
RUN apk add git
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/stripcall/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /multistage .
CMD ["./main"]