package twitter

import (
	"encoding/base64"
	"net/url"
	"os"
	"path"

	"github.com/ChimeraCoder/anaconda"
)

func PostPicture() error {
	anaconda.SetConsumerKey(os.Getenv("TWAPIKEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWAPIKEYSECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWACCESSTOKEN"), os.Getenv("TWACCESSTOKENSECRET"))

	exe, _ := os.Getwd()

	file, err := os.Open(path.Join(exe, "picture", "output", "output0.png"))
	if err != nil {
		return err
	}
	defer file.Close()

	f, _ := file.Stat()
	data := make([]byte, f.Size())
	file.Read(data)

	media, _ := api.UploadMedia(base64.StdEncoding.EncodeToString(data))

	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)

	_, err = api.PostTweet("", v)
	if err != nil {
		return err
	}

	return nil

}
