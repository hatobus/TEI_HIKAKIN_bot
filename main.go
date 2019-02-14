package main

import (
	"bufio"
	"log"
	"os"

	"github.com/hatobus/tei_hikakin/twitter"
	"github.com/hatobus/tei_hikakin/util"
	"github.com/hatobus/tei_hikakin/youtube"
)

func main() {
	util.Loadenv()

	channelelem := []string{"", "_TV", "_GAMES", "_BLOG"}

	isNew := false
	var channel, videoID string

	for _, elem := range channelelem {
		videoID, err := youtube.GetLatestMovieID(os.Getenv("HIKAKIN" + elem))
		if err != nil {
			log.Println(err)
		} else {
			log.Println(videoID)
		}

		// Checking videoID for new videos.
		f, err := os.Open("./HIKAKIN" + elem + ".txt")
		if err != nil {
			log.Fatalln(err)
		}

		scanner := bufio.NewScanner(f)

		var prevID string
		for scanner.Scan() {
			prevID = scanner.Text()
			log.Println(prevID)
		}

		if prevID != videoID {
			isNew = true
			channel = elem
			f.Close()
			break
		}

		f.Close()
	}

	if !isNew {
		log.Println("Already got thumbnail")
		// Golang's "os.Exist" can't execute defer method.
		os.Exit(0)
	}

	f, err := os.Create("./HIKAKIN" + channel + ".txt")
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

	err = twitter.PostPicture()
	if err != nil {
		log.Println(err)
	}

}
