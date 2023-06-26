FROM golang:1.20.5-alpine

WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o a.out server.go

EXPOSE 8080

CMD ["/a.out"]

