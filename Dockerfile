# https://hub.docker.com/_/golang
FROM golang:1.20

ADD . /app
WORKDIR /app

# Pre-copy/cache go.mod for pre-downloading dependencies and 
# only redownloading them in subsequent builds if they change.
RUN go mod download && go mod verify

RUN go build

CMD ["./run.sh"]
