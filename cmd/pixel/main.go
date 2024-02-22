package main

import (
	"errors"
	"flag"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nfnt/resize"
	"github.com/sachinsenal0x64/pixel"
)

func main() {
	var (
		x, y, w, h int
		url        string
	)

	flag.IntVar(&x, "x", 0, "")
	flag.IntVar(&y, "y", 0, "")
	flag.IntVar(&w, "w", 0, "")
	flag.IntVar(&h, "h", 0, "")
	flag.StringVar(&url, "url", "", "URL of the image")
	flag.Parse()

	log.SetPrefix("")

	var img image.Image
	var err error

	if url != "" {
		img, err = loadImageFromURL(url)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		img, _, err = image.Decode(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if w != 0 || h != 0 {
		img = resize.Thumbnail(
			uint(w), uint(h), img, resize.Bilinear,
		)
	}

	i, err := pixel.NewImage(img, x, y)
	if err != nil {
		log.Fatalln(err)
	}

	i.Show()

	defer i.Destroy()

	select {}
}

func loadImageFromURL(url string) (image.Image, error) {
	var reader io.Reader

	if strings.HasPrefix(url, "http") {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		reader = resp.Body
	} else if file, err := os.Open(url); err == nil {
		defer file.Close()
		reader = file
	} else {
		return nil, errors.New("unsupported URL format")
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return img, nil
}
