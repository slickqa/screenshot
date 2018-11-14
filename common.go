package screenshot

import "image"

type ScreenshotUtil interface {
	ScreenRect() (image.Rectangle, error)
	CaptureScreen() (*image.RGBA, error)
	CaptureRect(rect image.Rectangle) (*image.RGBA, error)
	Close()
}
