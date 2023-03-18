#stage 1
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o app ./cmd/app/app.go

#stage 2
FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]