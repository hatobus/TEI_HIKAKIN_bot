package youtube

import (
	"io"
	"os"
	"path"
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

	if len(landmark) == 0 {
		f.Seek(0, 0)
		outfile, err := os.Create(path.Join(exe, "picture", "output", "output0.png"))
		if err != nil {
			return err
		}
		defer outfile.Close()

		_, err = io.Copy(outfile, f)
		if err != nil {
			return err
		}

		return nil
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
