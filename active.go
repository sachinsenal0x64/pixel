package pixel

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
	"os"
)

const (
	atomActiveWindow = "_NET_ACTIVE_WINDOW"
)

func getFocusedWindow(X *xgbutil.XUtil) (xproto.Window, error) {
	// Get the currently focused window ID
	activeWindow, err := ewmh.ActiveWindowGet(X)
	fmt.Println(activeWindow)
	if err != nil {
		return 0, err
	}
	return activeWindow, nil
}

func isInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func getWMClass(X *xgbutil.XUtil, windowID xproto.Window) (string, string, error) {
	// Get the WM_CLASS property
	class, err := icccm.WmClassGet(X, windowID)
	if err != nil {
		return "", "", err
	}

	return class.Class, class.Instance, nil
}

func getActiveWindow(root xproto.Window) (xproto.Window, error) {

	X, err := xgbutil.NewConn()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer X.Conn().Close()

	i := xproto.InternAtom(
		x, true,
		uint16(len(atomActiveWindow)),
		atomActiveWindow,
	)

	a, err := i.Reply()
	if err != nil {
		return 0, err
	}

	// https://github.com/BurntSushi/xgb/blob/master/examples/get-active-window/main.go#L44
	c := xproto.GetProperty(
		x, false,
		root, a.Atom,
		xproto.GetPropertyTypeAny,
		0, (1<<32)-1,
	)

	r, err := c.Reply()
	if err != nil {
		panic(err)
	}

	// Get the currently focused window
	focusedWindow, err := getFocusedWindow(X)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if focusedWindow == 0 {
		fmt.Println("No focused window found.")
		os.Exit(3)
	}

	// Get window manager class of the focused window
	className, instanceName, err := getWMClass(X, focusedWindow)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	fmt.Println(className, instanceName)

	classNames := []string{"konsole", "org.wezfurlong.wezterm", "Alacritty", "kitty"}

	if isInSlice(className, classNames) {
		fmt.Println(focusedWindow)
		return xproto.Window(xgb.Get32(r.Value)), nil
	}
	return 0, errors.New("Terminal Not Support")
}
