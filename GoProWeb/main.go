package main

import(
	"net/http"
	"html/template"
	"log"
	"fmt"
)

const StaticDir string = "static/"

type Context struct{
	Title string
	StaticDir string
}


func vueBootstrap(w http.ResponseWriter,req *http.Request){

	context := &Context{Title:"vueBootsteap",StaticDir:StaticDir}
	render(w,"vuejs_bootstrap_demo",context)
}

func home(w http.ResponseWriter,req *http.Request){

	context := &Context{Title:"Homepage",StaticDir:StaticDir}
	render(w,"index",context)
}


func about(w http.ResponseWriter,req *http.Request){

	context := &Context{Title:"About",StaticDir:StaticDir}
	render(w,"about",context)
}

func base(w http.ResponseWriter,req *http.Request){

	context := &Context{Title:"Base",StaticDir:StaticDir}
	render(w,"base",context)
}



func render(w http.ResponseWriter,tmpl string,data interface{}){

	var templates = template.Must(template.ParseFiles("static/header.html","static/footer.html"))

	tmpl_list := []string{fmt.Sprintf("static/%s.html", tmpl)}
	t, err := templates.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.ExecuteTemplate(w,tmpl,data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


}


func main(){

	http.HandleFunc("/",home)
	http.HandleFunc("/about",about)
	http.HandleFunc("/base",base)
	http.HandleFunc("/vueBootstrap",vueBootstrap)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))

	err := http.ListenAndServe(":8016", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

