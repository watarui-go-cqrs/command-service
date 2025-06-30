FROM golang:1.24-alpine3.22
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/command
WORKDIR /go/src/command
COPY . /go/src/command
RUN go mod download
EXPOSE 8082
CMD ["go", "run", "cmd/server/main.go"]
