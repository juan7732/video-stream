# Vision Stream 

Going with go this time

## Usage

### Build the image

```bash
docker build -t go-rpi-camera-stream .
```

### Run the container

```bash
docker run -d --name go-rpi-camera-stream --restart always -p 8080:8080 go-rpi-camera-stream
```

### Access the stream

Open a web browser and go to `http://<raspberry-pi-ip>:8000/`

