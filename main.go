package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

type pixel struct {
	R int
	G int
	B int
}

// getting width and height of image
func getWidthAndHeight(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	// Close image after all execution
	defer file.Close()

	// get width and height of image
	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

func getAndStorePixels(file io.Reader, widthImage int, heightImage int) ([][]pixel, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	width, height := widthImage, heightImage
	var pixels [][]pixel
	for y := 0; y < height; y++ {
		var row []pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

// get rgba values for pixels
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) pixel {
	return pixel{int(r / 257), int(g / 257), int(b / 257)}
}

func main() {
	// get path of image via command line
	flagImagePath := flag.String("path", "none", "Select path of image you want to convert.")
	flag.Parse()

	// getting width and height
	width, height := getWidthAndHeight(*flagImagePath)

	// getting pixels
	file, err := os.Open(*flagImagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	pixels, err := getAndStorePixels(file, width, height)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	fmt.Println(pixels)
}
