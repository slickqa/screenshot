package screenshot

import (
	"fmt"
	"image"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type xScreenshotUtil struct {
	conn *xgb.Conn
}

func CreateScreenshotUtility() (ScreenshotUtil, error) {
	c, err := xgb.NewConn()
	if err != nil {
		return nil, fmt.Errorf("error connecting to X: %s", err.Error())
	}
	return &xScreenshotUtil{conn: c}, nil
}

func (s *xScreenshotUtil) Close() {
	s.conn.Close()
}

func (s *xScreenshotUtil) ScreenRect() (image.Rectangle, error) {
	screen := xproto.Setup(s.conn).DefaultScreen(s.conn)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	return image.Rect(0, 0, int(x), int(y)), nil
}

func (s *xScreenshotUtil) CaptureScreen() (*image.RGBA, error) {
	r := ScreenRect()
	return CaptureRect(r)
}

func (s *xScreenshotUtil) CaptureRect(rect image.Rectangle) (*image.RGBA, error) {
	screen := xproto.Setup(s.conn).DefaultScreen(s.conn)
	x, y := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(s.conn, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * x, image.Rect(0, 0, x, y)}
	return img, nil
}
