package main

import (
	"image"
	"reflect"
	"testing"
)

func Test_generateImage(t *testing.T) {
	type args struct {
		b []byte
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
		{
			"generate",
			args{
				[]byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5},
				2,
				3,
			},
			&image.Gray{
				Pix:    []uint8{0, 1, 2, 3, 4, 5},
				Stride: 2,
				Rect:   image.Rect(0, 0, 2, 3),
			},
		},
		{
			"overflow",
			args{
				[]byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6},
				2,
				3,
			},
			&image.Gray{
				Pix:    []uint8{0, 1, 2, 3, 4, 5, 6},
				Stride: 2,
				Rect:   image.Rect(0, 0, 2, 3),
			},
		},
		{
			"pad",
			args{
				[]byte{0x0, 0x1, 0x2, 0x3, 0x4},
				2,
				3,
			},
			&image.Gray{
				Pix:    []uint8{0, 1, 2, 3, 4, 0},
				Stride: 2,
				Rect:   image.Rect(0, 0, 2, 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateImage(tt.args.b, tt.args.w, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
