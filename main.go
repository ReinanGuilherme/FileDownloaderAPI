package main

import (
	"fileDownloaderAPI/src/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/download-excel", utils.DownloadExcelHandler)
	http.HandleFunc("/download-text", utils.DownloadTextHandler)
	http.HandleFunc("/download-csv", utils.DownloadCSVHandler)
	http.HandleFunc("/download-pdf", utils.DownloadPDFHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
