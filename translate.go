package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"os"
)

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.USLetterWidth, pdf.USLetterHeight)
	canvas.Translate(pdf.USLetterWidth/2, pdf.USLetterHeight/2)
	text := new(pdf.Text)
	text.SetFont(pdf.Times, 480)
	r := float32(45)
	for i := 0; i < int(360/r); i++ {
		canvas.Rotate(r * 0.0174532925)
		text.Text("a")
		canvas.DrawText(text)
	}
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
