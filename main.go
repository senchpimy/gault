package main

import (
	"fmt"
	"net/http"

	//"os"
	//	"os/exec"
	"html/template"
	//	"fmt"
	//	"log"
)

type Todo struct {
    Title string
    Done  bool
}


type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func Discos(w http.ResponseWriter ,r *http.Request)  {
	Data:=FormaterDiskInfo(GetInfoSystem())
	t:=template.Must(template.ParseFiles("./discos.html"))
	t.Execute(w,Data)
}

func main() {
	port := ":3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/Discos", Discos)
	http.ListenAndServe(port, mux)

}
