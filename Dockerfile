FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum  ./

COPY go.work go.work.sum ./

COPY ./submodules ./submodules

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o app ./cmd/app/app.go

CMD ["./app"]