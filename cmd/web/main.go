package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// create leveled logs
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

// Middleware для установки правильных MIME типов
func (app *application) setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		app.infoLog.Printf("Запрос к файлу: %s, расширение: %s", r.URL.Path, ext)

		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
			app.infoLog.Println("Установлен Content-Type: text/css")
		case ".woff2":
			w.Header().Set("Content-Type", "font/woff2")
			app.infoLog.Println("Установлен Content-Type: font/woff2")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
			app.infoLog.Println("Установлен Content-Type: application/javascript")
		case ".html":
			w.Header().Set("Content-Type", "text/html")
			app.infoLog.Println("Установлен Content-Type: text/html")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// set flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	// parse flags
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Применяем middleware к статическим файлам
	mux.Handle("/static/", app.setContentType(http.StripPrefix("/static", fileServer)))

	mux.HandleFunc("/", app.home)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
