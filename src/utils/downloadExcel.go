package utils

import (
	"log"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func DownloadExcelHandler(w http.ResponseWriter, r *http.Request) {
	texts := []string{"Texto 1", "Texto 2", "Texto 3"} // Slice de strings com os textos a serem inseridos no arquivo Excel

	file := excelize.NewFile() // Cria um novo arquivo Excel usando a biblioteca excelize
	sheetName := "Sheet1"      // Nome da planilha no arquivo Excel

	// Percorre o slice de strings e insere cada texto em uma célula da planilha
	for i, text := range texts {
		cellName := "A" + strconv.Itoa(i+1)          // Define o nome da célula com base no índice
		file.SetCellValue(sheetName, cellName, text) // Insere o texto na célula correspondente
	}

	fileName := "excelFile.xlsx" // Nome do arquivo para download

	// Define os cabeçalhos de resposta para indicar que é um arquivo para download
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	err := file.Write(w) // Escreve o arquivo Excel na resposta HTTP
	if err != nil {
		log.Println("Failed to write Excel file:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
