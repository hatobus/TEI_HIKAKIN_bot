package youtube

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/hatobus/tei_hikakin/models"
)

func GetLatestMovieID() (string, error) {
	URL := "https://www.googleapis.com/youtube/v3/search"
	YOUTUBEKEY := os.Getenv("YOUTUBEKEY")

	u, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	query := url.Values{
		"part":      {"id"},
		"channelId": {os.Getenv("HIKAKIN_CHANNELID")},
		"order":     {"date"},
		"key":       {YOUTUBEKEY},
	}

	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var channelinfo models.YoutubeResponse
	err = json.Unmarshal(body, &channelinfo)
	if err != nil {
		return "", err
	}

	return channelinfo.Items[0].ID.VideoID, nil

}
