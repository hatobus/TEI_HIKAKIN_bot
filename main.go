package main

import (
	"log"

	"github.com/hatobus/tei_hikakin/util"
	"github.com/hatobus/tei_hikakin/youtube"
)

func main() {
	util.Loadenv()

	videoID, err := youtube.GetLatestMovieID()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(videoID)
	}

	err = youtube.GetThumbnail(videoID)
	if err != nil {
		log.Println(err)
	}
}
