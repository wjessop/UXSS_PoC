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
	homeTempl     *template.Template
	template_vars TemplateVars
)

type TemplateVars struct {
	Addr      string
	TargetURL string
	UseSleep  string
}

func serveHTTP() {
	template_vars = TemplateVars{
		Addr:      *addr,
		TargetURL: *target_url,
		UseSleep:  *use_sleep,
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

		if *template_file != "" {
			if _, err := os.Stat(*template_file); os.IsNotExist(err) {
				panic("Custom template file not found")
			}

			homeTempl = template.Must(template.ParseFiles(*template_file))
		} else {
			bytes, err := Asset("data/templates/home.html")
			if err != nil {
				panic(err)
			}
			homeTempl = template.Must(template.New("home").Parse(string(bytes)))
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
		log.Println("Redirecting to", *non_target_url)
		http.Redirect(w, req, *non_target_url, 302)
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
