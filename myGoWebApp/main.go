package main

import (
    "html/template"
    "net/http"
    "log"
)

//Compile templates on start
var templates=template.Must(template.ParseFiles("header.html", "footer.html", "main.html", "about.html"))

//A Page structure
type Page struct {
    Title string
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
    templates.ExecuteTemplate(w, tmpl, data)
}

//The handlers.
func mainHandler(w http.ResponseWriter, r *http.Request) {
    display(w, "main", &Page{Title: "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    display(w, "about", &Page{Title: "About"})
}

func main(){
    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/about", aboutHandler)
    http.Handle("/bootstrap-3.3.7/", http.StripPrefix("/bootstrap-3.3.7/", http.FileServer(http.Dir("./bootstrap-3.3.7"))))
    //Listen on port 8080
    log.Fatal(http.ListenAndServe(":8020", nil))
}