FROM golang:1.17-buster

RUN apt-get update && \
    apt-get install -y v4l-utils libv4l-dev && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main

CMD ["./main"]
