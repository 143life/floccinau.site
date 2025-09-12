package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// create leveled logs
var infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

// Middleware для установки правильных MIME типов
func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		infoLog.Printf("Запрос к файлу: %s, расширение: %s", r.URL.Path, ext)

		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
			infoLog.Println("Установлен Content-Type: text/css")
		case ".woff2":
			w.Header().Set("Content-Type", "font/woff2")
			infoLog.Println("Установлен Content-Type: font/woff2")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
			infoLog.Println("Установлен Content-Type: application/javascript")
		case ".html":
			w.Header().Set("Content-Type", "text/html")
			infoLog.Println("Установлен Content-Type: text/html")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// set flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	// parse flags
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Применяем middleware к статическим файлам
	mux.Handle("/static/", setContentType(http.StripPrefix("/static", fileServer)))

	mux.HandleFunc("/", home)

	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
