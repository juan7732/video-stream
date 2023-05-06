# Vision Stream flask server for Raspberry Pi Docker Image

Simple flask server that streams the Raspberry Pi camera to a web browser.

## Usage

### Build the image

```bash
docker build -t vision-stream .
```

### Run the container

```bash
docker run -d --name video_streaming_app --restart always --device /dev/video0:/dev/video0 -p 8000:8000 vision-stream
```

### Access the stream

Open a web browser and go to `http://<raspberry-pi-ip>:8000/`

