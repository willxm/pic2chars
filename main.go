package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
	"strings"
)

var str1 string = `@@WW##$$XXoo**""==::''..--  `
var str2 string = `01. `

var chars []rune = []rune(str1)

func getChars(r, g, b, a uint32) string {
	if a == 0 {
		return ""
	}
	gray := int(0.2126*float64(r>>8) + 0.7152*float64(g>>8) + 0.0722*float64(b>>8))
	return string(chars[int(float64(gray)/256.0*float64(len(chars)))])
}

func main() {

	var zoom int
	var imgPath string
	// var type int

	flag.IntVar(&zoom, "z", 1, "zoom")
	flag.StringVar(&imgPath, "i", "test.jpeg", "image path")
	// flag.IntVar(&type, "t", 0, "chars type")

	flag.Parse()

	hz := 2 * zoom
	wz := zoom

	file, err := os.Open("input/" + imgPath)
	if err != nil {
		panic(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	var out string

	fmt.Println(img.Bounds().Min, img.Bounds().Max)
	for i := img.Bounds().Min.Y; i < img.Bounds().Max.Y; i = i + hz {
		for j := img.Bounds().Min.X; j < img.Bounds().Max.X; j = j + wz {
			fmt.Print(getChars(img.At(j, i).RGBA()))
			out += getChars(img.At(j, i).RGBA())
		}
		fmt.Print("\n")
		out += "\n"
	}

	f, err := os.Create("output/" + strings.Split(imgPath, ".")[0] + ".txt")
	if err != nil {
		panic(err)
	}
	f.WriteString(out)
	f.Close()
}
