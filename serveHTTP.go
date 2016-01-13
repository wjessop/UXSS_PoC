package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var (
	homeTempl     = template.Must(template.ParseFiles("templates/home.html"))
	template_vars TemplateVars
)

type TemplateVars struct {
	Addr         string
	NonTargetURL string
}

func serveHTTP() {
	template_vars = TemplateVars{
		Addr:         *addr,
		NonTargetURL: *non_target_url,
	}

	mux := http.NewServeMux()
	mux.Handle("/", serveHome())
	mux.Handle("/redirect", serveRedirect())
	mux.Handle("/sleep", serveSleep())

	s := &http.Server{
		Addr:           *addr,
		Handler:        handlers.LoggingHandler(os.Stdout, handlers.CompressHandler(context.ClearHandler(mux))),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Serving HTTP on address", *addr)
	log.Fatal(s.ListenAndServe())
}

func serveHome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.Error(w, "Not found", 404)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		homeTempl.Execute(w, template_vars)
	})
}

func serveRedirect() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/redirect" {
			http.Error(w, "Not found", 404)
			return
		}
		log.Println("Redirecting to", *protected_url)
		http.Redirect(w, req, *protected_url, 302)
	})
}

func serveSleep() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/sleep" {
			http.Error(w, "Not found", 404)
			return
		}
		time.Sleep(1 * time.Second)
		w.WriteHeader(204)
	})
}
