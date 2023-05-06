package main

import (
	"bytes"
	"fmt"
	"github.com/blackjack/webcam"
	"io"
	"log"
	"net/http"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		log.Fatal(err)
	}

	defer cam.Close()

	format_desc := cam.GetSupportedFormats()
	var format webcam.PixelFormat
	for f, desc := range format_desc {
		if desc == "MJPG" {
			format = f
			break
		}
	}
	f, width, height, err := cam.SetImageFormat(format, 640, 480)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Selected format: %s (%dx%d)\n", format_desc[f], width, height)

	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	for {
		err = cam.StartStreaming()
		if err != nil {
			log.Fatal(err)
		}

		frame, err := cam.ReadFrame()
		if err != nil {
			log.Println("Error capturing frame:", err)
			continue
		}

		buf := bytes.NewBuffer(frame)
		w.Write([]byte("\r\n--frame\r\nContent-Type: image/jpeg\r\n\r\n"))
		io.Copy(w, buf)

		cam.StopStreaming()
	}
}

func main() {
	http.HandleFunc("/", streamHandler)
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
