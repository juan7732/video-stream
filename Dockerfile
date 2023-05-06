FROM golang:1.17-buster

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN sudo apt-get update && sudo apt-get upgrade
RUN sudo apt-get install -y build-essential cmake git pkg-config libgtk-3-dev \
  libavcodec-dev libavformat-dev libswscale-dev libv4l-dev \
  libxvidcore-dev libx264-dev libjpeg-dev libpng-dev libtiff-dev \
  gfortran openexr libatlas-base-dev python3-dev python3-numpy \
  libtbb2 libtbb-dev libdc1394-22-dev

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["/app/main"]
