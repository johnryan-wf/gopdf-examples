package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.USLetterWidth, pdf.USLetterHeight)
	canvas.Translate(pdf.USLetterWidth/2, pdf.USLetterHeight/2)

	f, err := os.Open("testdata/suzanne.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rect := &pdf.Rectangle{
		pdf.Point{-100, -100}, // min
		pdf.Point{100, 100}}   // max

	// draw an image
	canvas.DrawImage(img, *rect)

	canvas.Close()

	err = doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
