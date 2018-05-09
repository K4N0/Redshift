package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dataanalyst", dataanalyst)
	http.HandleFunc("/hud_responders", hud_responders)
	http.HandleFunc("/icqa", icqa)
	http.HandleFunc("/learning", learning)
	http.HandleFunc("/palletsearch", palletsearch)
	http.HandleFunc("/test", test)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func dataanalyst(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dataanalyst.gohtml", nil)
}

func hud_responders(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hud_responders.gohtml", nil)
}

func icqa(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "icqa.gohtml", nil)
}

func learning(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "learning.gohtml", nil)
}

func palletsearch(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "palletsearch.gohtml", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "test.gohtml", nil)
}
