package youtube

import (
	"os"
	"path/filepath"

	"github.com/hatobus/Teikyo/callapi"
	img "github.com/hatobus/Teikyo/imgprocessing"
)

func GenTeikyo() error {

	exe, _ := os.Getwd()
	picdir := filepath.Join(exe, "picture", "thumbnail", "thumbnail.jpg")

	f, err := os.Open(picdir)
	if err != nil {
		return err
	}

	defer f.Close()

	landmark, err := callapi.DetectFace(f)
	if err != nil {
		return err
	}

	mul := false

	if len(landmark) > 1 {
		mul = true
	}

	for i, Mark := range landmark {
		LM := Mark.ToLandmark()
		err := img.GenTeikyo(f, LM, mul, 0, i)
		if err != nil {
			return err
		}
	}

	return nil
}
