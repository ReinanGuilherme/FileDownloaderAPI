package utils

import (
	"log"
	"net/http"
	"strings"
)

func DownloadTextHandler(w http.ResponseWriter, r *http.Request) {
	texts := []string{"Texto 1", "Texto 2", "Texto 3"} // Slice de strings

	fileName := "txtFile.txt" // Nome do arquivo para download

	// Concatena as strings do slice em um único texto
	content := strings.Join(texts, "\n")

	// Define os cabeçalhos de resposta
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/plain")

	// Escreve o conteúdo do arquivo na resposta HTTP
	_, err := w.Write([]byte(content))
	if err != nil {
		log.Println("Failed to write file to response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
