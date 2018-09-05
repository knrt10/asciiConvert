package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func getWidthAndHeight(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

func main() {

	// get path of image via command line
	flagImagePath := flag.String("path", "none", "Select path of image you want to convert.")
	flag.Parse()
	width, height := getWidthAndHeight(*flagImagePath)
	fmt.Println("Width:", width, "Height:", height)
}
