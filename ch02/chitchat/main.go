package main

import (
	"net/http"
	"text/template"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signip", signup)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		public_tmlp_files := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		private_tmpl_files := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}
		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmlp_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmlp_files...))
		}

		templates.ExecuteTemplate(w, "layout", threads)
	}
}
