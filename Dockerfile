FROM golang:1.24.1-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o project-management-api ./cmd/main.go

EXPOSE 8080

ENV PORT=8080

CMD ["./project-management-api"]