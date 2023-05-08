package utils

import (
	"log"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

func DownloadPDFHandler(w http.ResponseWriter, r *http.Request) {
	texts := []string{"Texto 1", "Texto 2", "Texto 3"} // Slice de strings

	fileName := "pdfFile.pdf" // Nome do arquivo para download

	// Cria um novo documento PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Adiciona uma nova página
	pdf.AddPage()

	// Define a fonte e o tamanho
	pdf.SetFont("Arial", "", 12)

	// Escreve o conteúdo do slice de strings no PDF
	for _, text := range texts {
		pdf.Cell(0, 10, text)
		pdf.Ln(10)
	}

	// Salva o PDF na resposta HTTP
	err := pdf.Output(w)
	if err != nil {
		log.Println("Failed to write PDF to response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Define os cabeçalhos de resposta
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/pdf")
}
