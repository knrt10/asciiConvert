package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/fatih/color"
)

type pixel struct {
	R float64
	G float64
	B float64
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
	return pixel{float64(r / 257), float64(g / 257), float64(b / 257)}
}

func formatMatrix(finalMatrix [][]string, width int, height int) {
	c := color.New(color.FgGreen).Add(color.Underline)
	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			finalMatrix[x][y] = finalMatrix[x][y] + finalMatrix[x][y] + finalMatrix[x][y]
		}
	}
	c.Println(finalMatrix[0])
}

func main() {

	asciiChars := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	maxPixelVal := 255

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
	// fmt.Println(pixels)

	// Converting the RGB tuples of pixels into single brightness numbers
	var brightness [][]int
	for y := 0; y < len(pixels); y++ {
		var row []int
		for _, v := range pixels[y] {
			var sum = 0
			sum = int(0.21*v.R + 0.72*v.G + 0.07*v.B)
			row = append(row, sum)
		}
		brightness = append(brightness, row)
	}

	// fmt.Println(brightness)

	// calculating max Value
	max := brightness[0][0]
	min := brightness[0][0]
	for _, value := range brightness {
		for _, k := range value {
			if k > max {
				max = k
			}
		}
	}

	// calculating min value
	for _, value := range brightness {
		for _, k := range value {
			if k < min {
				min = k
			}
		}
	}

	// fmt.Println(max, min)

	// Normalizing Matrix

	var normalize [][]int
	for y := 0; y < len(brightness); y++ {
		var row []int
		for _, v := range brightness[y] {
			r := int(maxPixelVal) * (v - min) / int(max-min)
			row = append(row, int(r))
		}
		normalize = append(normalize, row)
	}

	// Convert to ascii Characters
	var finalMatrix [][]string
	for y := 0; y < len(normalize); y++ {
		var row []string
		for _, v := range normalize[y] {
			row = append(row, string(asciiChars[int(v/(maxPixelVal)*len(asciiChars))]))
		}
		finalMatrix = append(finalMatrix, row)
	}

	formatMatrix(finalMatrix, width, height)

}
