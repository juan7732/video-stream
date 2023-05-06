package main

import (
	"fmt"
	"time"
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/korandiz/v4l"
)

const (
	cameraDevice = "/dev/video0"
	boundary     = "frame"
	listenAddr   = ":8080"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/stream", streamHandler)

	log.Printf("Starting server on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := `
		<html>
		<head><title>Go Pi Camera Stream</title></head>
		<body>
			<img src="/stream" style="width:640px;height:480px;" />
		</body>
		</html>
	`
	fmt.Fprintf(w, page)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	device, err := v4l.Open(cameraDevice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening video device: %v", err), http.StatusInternalServerError)
		return
	}
	defer device.Close()

	w.Header().Set("Content-Type", "multipart/x-mixed-replace;boundary="+boundary)
	multiPartWriter := multipart.NewWriter(w)
	defer multiPartWriter.Close()
	multiPartWriter.SetBoundary(boundary)

	for {
		frame, err := device.Capture()
		if err != nil {
			log.Printf("Error capturing frame: %v", err)
			continue
		}

		frameData := make([]byte, frame.Size())
		_, err = frame.Read(frameData)
		if err != nil {
			log.Printf("Error reading frame data: %v", err)
			continue
		}

		img, _, err := image.Decode(bytes.NewReader(frameData))
		if err != nil {
			log.Printf("Error decoding frame: %v", err)
			continue
		}

		partWriter, err := multiPartWriter.CreatePart(textproto.MIMEHeader{
			"Content-Type": {"image/jpeg"},
		})
		if err != nil {
			log.Printf("Error creating multipart writer: %v", err)
			continue
		}

		err = jpeg.Encode(partWriter, img, &jpeg.Options{Quality: 75})
		if err != nil {
			log.Printf("Error encoding frame: %v", err)
			continue
		}

		time.Sleep(33 * time.Millisecond)
	}
}
