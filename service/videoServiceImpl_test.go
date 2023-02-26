package service

import (
	"fmt"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	list, err := VideoServiceImpl{}.List(3)
	if err != nil {
		return
	}
	for _, video := range list {
		fmt.Println(video)
	}
}

func TestGetVideo(t *testing.T) {
	video, err := VideoServiceImpl{}.GetVideo(1, 2)
	if err != nil {
		return
	}
	fmt.Println(video)
}

func TestFeed(t *testing.T) {
	feed, t2, err := VideoServiceImpl{}.Feed(time.Now(), 2)
	if err != nil {
		return
	}
	for _, video := range feed {
		fmt.Println(video)
	}
	fmt.Println(t2)
}
