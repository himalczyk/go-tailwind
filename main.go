package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Render template
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/index.html")
        if err != nil {
            log.Fatal(err)
        }
        err = tmpl.Execute(w, nil)
        if err != nil {
            log.Fatal(err)
        }
    })

    // Start server
    log.Println("Listening on :8888")
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal(err)
    }
}
