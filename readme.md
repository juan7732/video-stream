# Vision Stream 

Going with go this time

## Usage

### Build the image

```bash
docker build -t video-stream-go .
```

### Run the container

```bash
docker run --rm --device=/dev/video0:/dev/video0 -p 8000:8000 video-stream-go
```

### Access the stream

Open a web browser and go to `http://<raspberry-pi-ip>:8000/`

