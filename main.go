package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackson-ajar/go-mjpeg"
	"github.com/korandiz/v4l"
)

func main() {
	camera, err := v4l.Open("/dev/video0")
	if err != nil {
		log.Fatalf("Error opening video device: %v", err)
	}

	defer camera.Close()

	stream := mjpeg.NewStream()

	go func() {
		for {
			frame, err := camera.Capture()
			if err != nil {
				log.Printf("Error capturing frame: %v", err)
			} else {
				stream.UpdateJPEG(frame)
			}
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	router.Handle("/stream", stream)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
