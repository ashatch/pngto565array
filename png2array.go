package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Printf("Please supply a PNG file as an argument\n\n")
		os.Exit(1)
	}

	fileName := os.Args[1]
	var file, err = os.Open(fileName)

	if err != nil {
		fmt.Printf("Could not open file\n\n")
		os.Exit(1)
	}

	var img, decodeErr = png.Decode(file)
	if decodeErr != nil {
		fmt.Printf("Could not decode PNG\n\n")
		os.Exit(1)
	}

	var size = (img.Bounds().Max.X) * (img.Bounds().Max.Y-1);
	fmt.Printf("# converted PNG as RGB 5/6/5 color format array for use in Adafruit::GFX drawRGBBitmap\n")
	fmt.Printf("const unsigned short bitmap[%d] PROGMEM={", size)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		fmt.Print("  ")
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := img.At(x, y)

			r, g, b, _ := c.RGBA()
			r5 := ((r >> 3) & 0x1f) << 11
			g6 := ((g >> 2) & 0x3f) << 5
			b5 := (b >> 3) & 0x1f
			rgb565 := r5 | g6 | b5

			fmt.Printf("0x%04x", rgb565)
			if !(x == img.Bounds().Max.X-1 && y == img.Bounds().Max.Y-1) {
				fmt.Printf(", ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Printf("};\n")
}
