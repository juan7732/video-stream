# Vision Stream 

Going with go this time

## Usage

### Build the image

```bash
docker build -t go-pi-camera .

```

### Run the container

```bash
docker run -d --name vision-stream --restart always --rm -p 8080:8080 --device /dev/video0:/dev/video0 go-pi-camera

```

### Access the stream

Open a web browser and go to `http://<raspberry-pi-ip>:8000/`

