package text2img

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"testing"
)

func TestText2img(t *testing.T) {

	f := func() {
		d, err := NewDrawer(Params{
			FontPath: "./miniHei.ttf",
			Width:    800,
			Height:   600,
		})
		log.Print(err)

		img, err := d.Draw("思念是一种病")
		log.Print(err)
		file, err := os.Create(fmt.Sprint(1, ".jpg"))
		log.Print(err)
		defer file.Close()
		log.Print(err)
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
		log.Print(err)
	}
	f()

}
