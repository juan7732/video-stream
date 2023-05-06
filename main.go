package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"gocv.io/x/gocv"
)

const (
	cameraDevice = 0
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
			<h1>Zora's Webstream</h1>
			<img src="/stream" style="width:640px;height:480px;" />
		</body>
		</html>
	`
	fmt.Fprintf(w, page)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	webcam, err := gocv.OpenVideoCapture(cameraDevice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening video device: %v", err), http.StatusInternalServerError)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	w.Header().Set("Content-Type", "multipart/x-mixed-replace;boundary="+boundary)
	multiPartWriter := multipart.NewWriter(w)
	defer multiPartWriter.Close()
	multiPartWriter.SetBoundary(boundary)

	for {
		if ok := webcam.Read(&img); !ok {
			log.Printf("Error reading frame")
			continue
		}
		if img.Empty() {
			continue
		}

		buf, err := gocv.IMEncode(".jpg", img)
		if err != nil {
			log.Printf("Error encoding frame: %v", err)
			continue
		}

		bufBytes := buf.GetBytes()
		defer buf.Close()

		partWriter, err := multiPartWriter.CreatePart(textproto.MIMEHeader{
			"Content-Type": {"image/jpeg"},
		})
		if err != nil {
			log.Printf("Error creating multipart writer: %v", err)
			continue
		}

		_, err = partWriter.Write(bufBytes)
		if err != nil {
			log.Printf("Error writing frame: %v", err)
			continue
		}
	}
}
