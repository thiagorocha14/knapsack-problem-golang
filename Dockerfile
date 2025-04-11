FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app

RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main