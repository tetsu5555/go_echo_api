FROM golang:1.12
ENV GO111MODULE=on

RUN apt-get update -qq && \
    apt-get install -y vim && \
    go get -u github.com/pilu/fresh

WORKDIR /go/src/api

COPY api/go.mod .
COPY api/go.sum .

RUN go mod download

COPY . .

CMD ["fresh"]