package p

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	model "github.com/hatobus/tei_hikakin/models"
	"github.com/hatobus/tei_hikakin/youtube"
)

func GenerateTeikyoEye(w http.ResponseWriter, r *http.Request) {
	if challenge := r.URL.Query().Get("hub.challenge"); challenge != "" {
		log.Println(challenge)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, challenge)
		return
	}

	sub := &model.Subsc{}

	if err := xml.NewDecoder(r.Body).Decode(sub); err != nil {
		log.Println(err)
		fmt.Fprint(w, "")
		return
	}

	if sub.Entry.Title == "" || sub.Entry.VideoId == "" {
		log.Println("title or video id is null")
		return
	} else if sub.Entry.VideoId == os.Getenv("LATEST_VIDEO_ID") {
		log.Println("already generated")
		return
	}

	err = youtube.CFGetThumbnail(sub.Entry.VideoId)
	if err != nil {
		log.Println("youtube getthumbnail error : %v", err)
		return
	}

	log.Printf("Video Title : %s, Video ID : %s", sub.Entry.Title, sub.Entry.VideoId)

	err = youtube.GenTeikyo()
	if err != nil {
		log.Println("youtube genteikyo error : %v", err)
		return
	}
}
