FROM golang:1.25.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-service ./main.go

EXPOSE 8080


CMD ["/go-service"]