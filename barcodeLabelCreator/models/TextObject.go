package models

// TextObject is a text type object of barcode label.
// TextObject later will be processed into image under 'ImageObject'
// along with its position.
type TextObject struct {
	Value                string // the text it displays
	FontSize             float64
	FontStyle            string  // only works if the font style is installed.
	FontWeight           string  // bold or normal
	XPosition, YPosition float64 // text's x and y position later on the barcode label.
}
