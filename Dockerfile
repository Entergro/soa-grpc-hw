FROM golang:1.17

WORKDIR /usr/src/server

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY ./server/ ./server/
COPY ./pkg/proto/hw/ ./pkg/proto/hw

RUN go build -v -o /usr/local/bin/server ./server/

CMD ["server"]
