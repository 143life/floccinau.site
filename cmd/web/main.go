package main

import (
	"log"
	"net/http"
	"path/filepath"
)

// Middleware для установки правильных MIME типов
func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		log.Printf("Запрос к файлу: %s, расширение: %s", r.URL.Path, ext)
		
		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
			log.Println("Установлен Content-Type: text/css")
		case ".woff2":
			w.Header().Set("Content-Type", "font/woff2")
			log.Println("Установлен Content-Type: font/woff2")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
			log.Println("Установлен Content-Type: application/javascript")
		case ".html":
			w.Header().Set("Content-Type", "text/html")
			log.Println("Установлен Content-Type: text/html")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	
	// Применяем middleware к статическим файлам
	mux.Handle("/static/", setContentType(http.StripPrefix("/static", fileServer)))

	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
