package main

import "fmt"

func main() {
	videos := getVideos() //return a slice of video
	for i, _ := range videos {
		//Update video description
		videos[i].Description = videos[i].Description + " Updated"
	}
	fmt.Println(videos)

	saveVideos(videos)
}
