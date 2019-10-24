package youtube

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/hatobus/tei_hikakin/models"
)

func GetLatestMovieID(channelID string) (string, error) {
	URL := "https://www.googleapis.com/youtube/v3/search"
	YOUTUBEKEY := os.Getenv("YOUTUBEKEY")

	u, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	query := url.Values{
		"part":      {"id"},
		"channelId": {channelID},
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

func GetThumbnail(videoID string) error {
	URL := "https://img.youtube.com/vi/"
	exe, _ := os.Getwd()
	savedir := filepath.Join(exe, "picture", "thumbnail", "thumbnail.jpg")

	u, err := url.Parse(URL)
	if err != nil {
		return nil
	}

	u.Path = path.Join(u.Path, videoID, "maxresdefault.jpg")

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	file, err := os.Create(savedir)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil

}

func CFGetThumbnail(videoID string) error {
	URL := "https://img.youtube.com/vi/"
	exe, _ := os.Getenv("ENTRY_POINT")
	savedir := filepath.Join(exe, "picture", "thumbnail", "thumbnail.jpg")

	u, err := url.Parse(URL)
	if err != nil {
		return nil
	}

	u.Path = path.Join(u.Path, videoID, "maxresdefault.jpg")

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	file, err := os.Create(savedir)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil

}
