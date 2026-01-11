# multistage docker build. This redices the size of the final docker image.
# stage 1 to build the app
FROM golang:alpine as builder 

WORKDIR /build 

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the api directory
COPY . .

# Build the application
RUN go build -o main .

# stage 2 deploys the app build in stage 1
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 3000
CMD ["./main"]