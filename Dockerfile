FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-proxy-backend

EXPOSE 8080

CMD [ "/go-proxy-backend" ]

