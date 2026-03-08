FROM golang:1.24.0-alpine3.21 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/microservices-project
COPY go.mod go.sum ./
COPY vendor vendor
COPY account account
RUN GO111MODULE=on CGO_ENABLED=0 go build -mod vendor -o /go/bin/app ./account/cmd/account

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/app /app
EXPOSE 8080
CMD ["/app"]