package cmd

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"reflect"

	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "", "path of your file for which you want to convert ASCII Art")
	rootCmd.Flags().IntVarP(&widthFlag, "width", "w", 100, "width of final file")
}

var asciiChar = "MND8OZ$7I?+=~:,.."

// function to return width and image
func getWidthAndImage(imagePath string, width int) (image.Image, int) {
	file, err := os.Open(imagePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	return img, width
}

// conversion to asciiArt

func asciiArt(img image.Image, w, h int) []byte {
	table := []byte(asciiChar)
	buffer := new(bytes.Buffer)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buffer.WriteByte(table[pos])
		}
		_ = buffer.WriteByte('\n')
	}
	return buffer.Bytes()
}

// get height of image

func getHeight(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	height := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(height), img, resize.Lanczos3)
	return img, w, height
}

var widthFlag int

var rootCmd = &cobra.Command{
	Use:   "asciiArt",
	Short: "AscII CLI",
	Long:  `This is a terminal client to create ASCII art from any image built with love by knrt10 in an effort to learn Go`,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		image, width := getWidthAndImage(path, int(widthFlag))
		finalASCIIArt := asciiArt(getHeight(image, width))
		fmt.Println(string(finalASCIIArt))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
