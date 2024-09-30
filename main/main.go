package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"piscine"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", ServeIndex)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		asciiArt := ""
		text := r.FormValue("text")
		if strings.TrimSpace(text) == "" {
			http.Error(w, "400 - Bad Request", http.StatusBadRequest)
			return
		}
		if !ISvalid(text) {
			http.Error(w, "400 - Bad Request", http.StatusBadRequest)
			return
		}
		banner := r.FormValue("banner")
		if banner == "" {
			banner = "standard" // Set default banner
		}
		banner += ".txt"
		in := piscine.Load(banner)
		if in == nil {
			http.Error(w, "400 - Bad Request", http.StatusInternalServerError)
			return
		}
		asciiArt = piscine.PrintOutput(in, text)
		tmpl, err := template.ParseFiles("template/index.html")
		if err != nil {
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := struct {
			Art string
		}{
			Art: asciiArt,
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	// Handle GET request to serve the form
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
	}
}

func ISvalid(s string) bool {
	flag := true
	for _, ch := range s {
		if ch >= 32 && ch <= 126 {
			continue
		}
		if ch == '\n' || ch == '\r' {
			continue
		}
		flag = false
	}
	return flag
}
