FROM golang:1.18 as builder

RUN apt update

ENV GO111MODULE on
WORKDIR /var/www/src

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY ./src /var/www

COPY ["./go.mod", "./go.sum", ".air.toml", "/var/www/"]

# RUN go mod download

CMD go run main.go

FROM golang:1.18 as dev

RUN apt update

ENV GO111MODULE on
WORKDIR /var/www/src

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install golang.org/x/lint/golint@latest && \
  go install golang.org/x/tools/gopls@latest && \
  go install honnef.co/go/tools/cmd/staticcheck@latest && \
  go install github.com/cosmtrek/air@latest

COPY ./src /var/www
