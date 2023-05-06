FROM golang:1.17-buster

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PKG_CONFIG_PATH /var/lib/docker/overlay2/kj3mfvsc6428qoj182qa9ctze/diff/usr/lib/arm-linux-gnueabihf/pkgconfig/opencv4.pc:$PKG_CONFIG_PATH

RUN go build -o main .

CMD ["/app/main"]
