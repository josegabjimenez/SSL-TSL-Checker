FROM golang:1.25-alpine

WORKDIR /app

# Install git (required for go install)
RUN apk add --no-cache git

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

COPY go.mod ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
