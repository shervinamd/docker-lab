FROM golang:1.21-alpine3.18 AS builder
WORKDIR /usr/src/goapp
COPY app.go .
RUN go build -v -o ./app ./app.go

FROM alpine:3.18
COPY --from=builder /usr/src/goapp/app /usr/local/bin/
EXPOSE 8080
CMD /usr/local/bin/app
