package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"143life.floccinau.site/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

// create leveled logs
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	draft    *models.DraftModel
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
	dsn := flag.String("dsn", "web:pass@/flocci?parseTime=true", "MySQL data source name")
	// parse flags
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create connection pool to MySQL
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		draft:    &models.DraftModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
