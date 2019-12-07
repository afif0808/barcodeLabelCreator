package controllers

import "net/http"

// CreateAndPrintBarcodeLabel create and print barcode label.
// It receives input from client which defines elements of a barcode label and its setting
// there are three types of barcode label element: text , barcode and image.
func CreateAndPrintBarcodeLabel() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
