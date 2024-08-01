package entity

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GenerateCertificate(name string) string {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Certificate of Completion")
	pdf.Ln(20)
	pdf.Cell(40, 10, fmt.Sprintf("This certifies that %s", name))
	pdf.Ln(20)
	pdf.Cell(40, 10, "has successfully completed the course.")

	fileName := fmt.Sprintf("%s_certificate.pdf", name)
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return fileName
}
