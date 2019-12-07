package models

import "image"

type CreateTextImage func(TextObject) (ImageObject, error)
type CreateBarcode func(BarcodeObject) (ImageObject, error)
type CreateBarcodeLabel func([]ImageObject) (image.Image, error)
type PrintBarcodeLabel func(img image.Image, qty int) error
