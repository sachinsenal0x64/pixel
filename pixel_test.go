package pixel

import (
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 9706b70 (Lot of Bug Fixes)
	"image"
	_ "image/jpeg"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

<<<<<<< HEAD
=======
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
>>>>>>> 9706b70 (Lot of Bug Fixes)
	"github.com/nfnt/resize"
)

func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, content)
	if err != nil {
		return err
	}

	return nil
}

func TestImage(t *testing.T) {
<<<<<<< HEAD
	img1, err := getImage("https://investnews.com.br/wp-content/uploads/2021/06/doge-1200x800.jpg")
=======
	X, err := xgbutil.NewConn()
>>>>>>> 9706b70 (Lot of Bug Fixes)
	if err != nil {
		t.Fatal(err)
	}

	defer X.Conn().Close()

	activeWindow, err := ewmh.ActiveWindowGet(X)

	// Overwrite the contents of "output.txt" with the new active window ID
	err = os.WriteFile("output.txt", []byte(fmt.Sprintf("%d", activeWindow)), 0644)
	if err != nil {
		t.Fatal(err)
	}

	img1, err := getImage("https://investnews.com.br/wp-content/uploads/2021/06/doge-1200x800.jpg")
	if err != nil {
		t.Fatal(err)
	}

	img2, err := getImage("_testdata/go.jpg")
	if err != nil {
		t.Fatal(err)
	}

	j, err := NewImage(img2, 300, 50)
	if err != nil {
		t.Fatal(err)
	}

	i, err := NewImage(img1, 700, 50)
	if err != nil {
		t.Fatal(err)
	}

	defer i.Clear()

	time.Sleep(time.Second * 8)

	defer i.Destroy()

	defer j.Clear()
	time.Sleep(time.Second * 8)
	defer j.Destroy()

}

func getImage(url string) (image.Image, error) {
	var reader io.Reader

	if strings.HasPrefix(url, "http") {
		r, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer r.Body.Close()

		reader = r.Body
	} else {
		f, err := os.Open(url)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		reader = f
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	img = resize.Thumbnail(
		300, 300,
		img,
		resize.Bilinear,
	)

	return img, nil
}
