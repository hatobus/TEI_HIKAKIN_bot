package main

import (
	"bufio"
	"log"
	"os"

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

	// Checking videoID for new videos.
	f, err := os.Open("./videoinfo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)

	var prevID string
	for scanner.Scan() {
		prevID = scanner.Text()
		log.Println(prevID)
	}

	if prevID == videoID {
		log.Println("Already got thumbnail")
		// Golang's "os.Exist" can't execute defer method.
		f.Close()
		os.Exit(0)
	}

	f.Close()

	f, err = os.Create("./videoinfo.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if _, err = f.WriteString(videoID); err != nil {
		log.Fatalln(err)
	}

	err = youtube.GetThumbnail(videoID)
	if err != nil {
		log.Println(err)
	}

	err = youtube.GenTeikyo()
	if err != nil {
		log.Println(err)
	}
}
