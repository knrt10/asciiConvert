package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"time"

	"os"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
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
		fmt.Println(err)
	}

	defer file.Close()
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
var path string

var rootCmd = &cobra.Command{
	Use:   "asciiConvert",
	Short: "AscII CLI",
	Long:  `This is a terminal client to create ASCII art from any image, built with love by knrt10 in an effort to learn Go`,
	Run: func(cmd *cobra.Command, args []string) {
		display(path, int(widthFlag))
	},
}

func display(path string, width int) {

	_, extension, err := mimetype.DetectFile(path)
	if err != nil {
		fmt.Println("Error finding type of file,\n Error:", err.Error())
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Couldn't open file,\n Error:", err.Error())
		os.Exit(1)
	}

	if extension == "gif" {
		displayGif(file, width)
	} else {
		displayImage(file, width)
	}
}

func displayImage(file *os.File, width int) {

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Couldn't decode image,\nError:", err.Error())
		os.Exit(1)
	}

	finalASCIIArt := asciiArt(getHeight(img, width))
	fmt.Println(string(finalASCIIArt))
}

func displayGif(file *os.File, width int) {

	gifImg, err := gif.DecodeAll(bufio.NewReader(file))
	if err != nil {
		fmt.Println("Couldn't decode image,\nError:", err.Error())
		os.Exit(1)
	}

	loopCount := 0
	for {
		for i, frame := range gifImg.Image {
			frameImage := frame.SubImage(frame.Rect)
			// To clear the shell
			fmt.Println("\033[2J")
			fmt.Printf("%s", asciiArt(getHeight(frameImage, width)))
			time.Sleep(time.Duration((time.Second * time.Duration(gifImg.Delay[i])) / 100))
		}
		// If gif is infinite loop
		if gifImg.LoopCount == 0 {
			continue
		}
		loopCount++
		if loopCount == gifImg.LoopCount {
			break
		}
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
