package models

import "image"

//ImageObject contains an image and its position on barcode label output
type ImageObject struct {
	image.Image
	XPosition, YPosition float64
}
