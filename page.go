package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"os"
)

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.USLetterWidth, pdf.USLetterHeight)
	canvas.Translate(100, pdf.USLetterHeight-100)
	text := new(pdf.Text)
	text.SetFont(pdf.Times, 12)
	text.Text("hello world")
	canvas.DrawText(text)
	canvas.Close()

	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
