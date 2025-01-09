package public

import (
	"log"
	"net/http"
	"os"
)

func Main_Page(w http.ResponseWriter, r *http.Request) {
	filePath := "./../frontend/index.html"
	if _, err := os.Stat(filePath); err != nil {
		http.Error(w, "Archivo no encontrado", http.StatusNotFound)
		log.Printf("Error: %v\n", err)
		return
	}
	http.ServeFile(w, r, filePath)
}