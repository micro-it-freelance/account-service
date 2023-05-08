FROM golang:alpine AS builder

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

COPY ./go.work ./
COPY ./go.work.sum ./

COPY ./submodules/core/go.mod ./submodules/core/go.mod
COPY ./submodules/core/go.sum ./submodules/core/go.sum 

COPY ./submodules/protoc/go.mod ./submodules/protoc/go.mod
COPY ./submodules/protoc/go.sum ./submodules/protoc/go.sum

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o app ./cmd/app/app.go

CMD ["./app"]