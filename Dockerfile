FROM golang:latest

LABEL maintainer="Aram Ceballos <aramgonzalez12@hotmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build api/app.go

CMD ["./app"]