package main

import (
	"encoding/json"
	"io/ioutil"
)

type Video struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Url         string `json:"url"`
}

func getVideos() (videos []Video) {
	fileBytes, err := ioutil.ReadFile("./video.json")
	if err != nil {
		panic(err)
	}
	//convert bytes to struct
	err = json.Unmarshal(fileBytes, &videos)
	if err != nil {
		panic(err)
	}
	return videos
}

func saveVideos(video []Video) {
	//convert struct to bytes
	videoBytes, err := json.Marshal(video)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
