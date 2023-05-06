FROM golang:1.17

WORKDIR /app

COPY . .

RUN go mod init video-stream-go
RUN go get -u github.com/blackjack/webcam
RUN go build -o main main.go

EXPOSE 8000

CMD ["./main"]