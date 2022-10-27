FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY . ./
RUN go mod download

RUN go build -o gateway-server

EXPOSE 8080
CMD ["./gateway-server"]