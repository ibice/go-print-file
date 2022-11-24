package main

import (
	"flag"
	"image/png"
	"log"
	"math"
	"os"
)

var (
	out    = flag.String("o", "", "output file")
	wRatio = flag.Int("w-ratio", 1, "width ratio")
	hRatio = flag.Int("h-ratio", 1, "heigth ratio")
	width  = flag.Int("w", 0, "width in pixels")
	height = flag.Int("h", 0, "heigth in pixels")
	wGrow  = flag.Int("grow-w", 0, "grow width in pixels")
	hGrow  = flag.Int("grow-h", 0, "grow heigth in pixels")
	skip   = flag.Int("s", 0, "skip leading bytes")
	in     = flag.String("i", "", "input file, the executable itself by default")
)

func parseArgs() {
	flag.Parse()

	if *out == "" {
		log.Fatal("-o is required")
	}

	var err error
	if *in == "" {
		*in, err = os.Executable()
		if err != nil {
			log.Fatalln("Getting executable path:", err)
		}
	}
}

func sizeByRatio(b []byte) (int, int) {
	// With and heigth — floating point
	fh := math.Sqrt(float64(*wRatio) * float64(len(b)) / float64(*hRatio))
	fw := float64(*wRatio) * fh / float64(*hRatio)
	log.Printf("Floating point size: w=%.2f, h=%.2f", fw, fh)

	// Width and heigth in pixels
	h := int(math.Ceil(fh))
	w := int(math.Ceil(fw))
	h += h % 2
	w += w % 2

	return w, h
}

func sizeFromConstraints(b []byte) (int, int) {
	if *height != 0 && *width != 0 {
		return *width, *height
	}
	if *width != 0 {
		return *width, int(math.Ceil(float64(len(b)) / float64(*width)))
	}
	if *height != 0 {
		return int(math.Ceil(float64(len(b)) / float64(*height))), *height
	}
	panic("no dimension was specified")
}

func main() {
	parseArgs()

	b, err := os.ReadFile(*in)
	if err != nil {
		log.Fatalln("Reading executable:", err)
	}

	if *skip != 0 {
		if *skip >= len(b) {
			log.Fatalf("Cannot skip %d bytes: file is %d bytes long", *skip, len(b))
		}
		b = b[*skip:]
	}

	var w, h int
	if *height == 0 && *width == 0 {
		w, h = sizeByRatio(b)
	} else {
		w, h = sizeFromConstraints(b)
	}

	img := generateImage(b, w+*wGrow, h+*hGrow)

	log.Printf("Writing image to %s", *out)
	f, err := os.Create(*out)
	if err != nil {
		log.Fatalln("Opening output file:", err)
	}

	if err = png.Encode(f, img); err != nil {
		log.Fatalln("Encoding image:", err)
	}

	log.Print("Done")
}
