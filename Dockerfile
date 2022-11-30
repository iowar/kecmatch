FROM golang:1.19

WORKDIR /usr/src/app
COPY . .

RUN go mod download
RUN go mod verify

WORKDIR /usr/src/app/build
RUN go build -ldflags "-s -w" -o /usr/local/bin/kecmatch

ENTRYPOINT ["/usr/local/bin/kecmatch"]
