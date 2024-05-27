FROM golang:1.22.3-bullseye

WORKDIR /workdir

ENV GO111MODULE=on \
    GOBIN=/usr/local/bin

COPY go.mod go.sum ./

RUN go install github.com/bufbuild/buf/cmd/buf@v1.32.1
RUN apt update && apt install -y protobuf-compiler
RUN go mod download

CMD ["buf", "dep", "update"]

# docker build -t buf-error-example .
# docker run -v ${PWD}:/workdir buf-error-example
