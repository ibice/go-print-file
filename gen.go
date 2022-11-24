package main

import (
	"image"
	"log"
)

func generateImage(b []byte, w, h int) *image.Gray {
	log.Printf("Generating image with len(b)=%d, w=%d, h=%d", len(b), w, h)
	win := w * h
	if len(b) < win {
		pad := win - len(b)
		log.Printf("Pad size: %d", pad)
		b = append(b, make([]byte, pad)...)
	}
	return &image.Gray{
		Pix:    b,
		Stride: w,
		Rect: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: w, Y: h},
		},
	}
}
