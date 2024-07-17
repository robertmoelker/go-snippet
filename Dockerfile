FROM golang:1.22.5-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

ENTRYPOINT ["air", "-c", ".air.toml"]
