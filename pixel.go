package pixel

import (
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
	"image"
)

// Image is the structure for the image
type Image struct {
	img *xgraphics.Image
	win *xwindow.Window
}

// NewImage makes a new image. This also initializes X if
// it's not initialized.

func NewImage(img image.Image, X, Y int) (*Image, error) {
	if x == nil {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

	bounds := img.Bounds()

	w, err := newChildWindow(
		X, Y,
		bounds.Dx(),
		bounds.Dy(),
	)

	if err != nil {
		return nil, err
	}

	// Make a new Image
	i := &Image{
		img: xgraphics.NewConvert(xutil, img),
		win: w,
	}

	if err := i.img.XSurfaceSet(i.win.Id); err != nil {
		return nil, err
	}

	if err := i.img.XDrawChecked(); err != nil {
		return nil, err
	}

	i.img.XPaint(i.win.Id)

	i.Show()

	return i, nil
}

// Show shows the image
func (i *Image) Show() {
	i.win.Map()
}

// Clear clears the image
func (i *Image) Clear() {
	i.win.Unmap()
}

// Destroy destroys the image and window, freeing up
// resources

func (i *Image) Destroy() {
	i.Clear()
	i.img.Destroy()
	i.win.Destroy()
}
