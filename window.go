package pixel

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xwindow"
)

func newParentWindow() *xwindow.Window {
	w, err := xwindow.Create(xutil, parent.Id)
	if err != nil {
		panic(err)
	}

	return w
}

func newChildWindow(X, Y, width, height int) (*xwindow.Window, error) {
	w, err := xwindow.Generate(xutil)
	if err != nil {

		return nil, err
	}

	// Create the window at the root level, we'll
	// tell X to re-parent it later
	if err := w.CreateChecked(
		xutil.RootWin(),
		X, Y, width, height,

		// This tells the window manager to not
		// touch the window, including overriding
		// the parent
		xproto.CwOverrideRedirect, 1,
	); err != nil {
		return nil, err
	}

	// This reparents the child window to its proper
	// parent, which is the terminal

	if err := xproto.ReparentWindowChecked(
		x, w.Id, parent.Id, int16(X), int16(Y),
	).Check(); err != nil {
		return nil, err
	}

	return w, nil
}
