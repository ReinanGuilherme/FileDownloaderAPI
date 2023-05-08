package utils

import (
	"encoding/csv"
	"log"
	"net/http"
	"strings"
)

func DownloadCSVHandler(w http.ResponseWriter, r *http.Request) {
	texts := []string{"Texto 1", "Texto 2", "Texto 3"} // Slice de strings

	fileName := "csvFile.csv" // Nome do arquivo para download

	// Cria um buffer para armazenar o conteúdo CSV
	csvContent := &strings.Builder{}

	// Cria um escritor CSV usando o buffer
	writer := csv.NewWriter(csvContent)

	// Escreve cada texto no escritor CSV
	for _, text := range texts {
		err := writer.Write([]string{text})
		if err != nil {
			log.Println("Failed to write CSV content:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	// Encerra a gravação do CSV
	writer.Flush()

	// Verifica erros no escritor CSV
	if err := writer.Error(); err != nil {
		log.Println("Failed to write CSV content:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Define os cabeçalhos de resposta
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/csv")

	// Escreve o conteúdo do arquivo na resposta HTTP
	_, err := w.Write([]byte(csvContent.String()))
	if err != nil {
		log.Println("Failed to write file to response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
