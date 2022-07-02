package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/", Hello)
	http.HandleFunc("/v", HandleGetVideos)
	http.HandleFunc("/v/update", HandleUpdateVideos)

	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	for header, value := range r.Header {
		fmt.Printf("key: %v \t Value: %v \n", header, value)
	}

	w.Header().Add("TestHeader", "TestValue")

	w.Write(([]byte("Hello!")))
}

func HandleGetVideos(w http.ResponseWriter, r *http.Request) {
	videos := getVideos()

	videoBytes, err := json.Marshal(videos) 
	
	if err != nil {
		panic(err)
	}
	w.Write(videoBytes)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// update our videos here!
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		
		// validate the content to see that they are legitimate videos
		var videos []video
		err = json.Unmarshal(body, &videos)
		if err != nil {
			w.WriteHeader(400) // write to the request
			fmt.Fprintf(w, "Bad Request") // write to the response
		}

		saveVideos(videos)
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not support")
	}


	// w.Write()
}