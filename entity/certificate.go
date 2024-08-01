package entity

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func getImageSize(imagePath string) (float64, float64) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return 210, 297
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return 210, 297
	}

	const dpi = 300.0
	widthMm := float64(img.Width) * 25.4 / dpi
	heightMm := float64(img.Height) * 25.4 / dpi

	return widthMm, heightMm
}

func GenerateCertificate(name string, templatePath string) string {
	imageWidth, imageHeight := getImageSize(templatePath)

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	pdf.Image(templatePath, 0, 0, imageWidth, imageHeight, false, "", 0, "")

	fontSize := 25.0
	pdf.SetFont("Helvetica", "B", fontSize)
	uppercaseName := strings.ToUpper(name)

	text := uppercaseName
	textWidth := pdf.GetStringWidth(text)
	textHeight := fontSize * 0.35277778

	posX := (imageWidth - textWidth) / 2
	posY := (imageHeight - textHeight) / 2

	pdf.SetXY(posX, posY)
	pdf.Cell(textWidth, textHeight, text)

	fileName := fmt.Sprintf("%s_certificate.pdf", name)
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Println("Error saving PDF:", err)
	}

	return fileName
}
