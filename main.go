package main

import (
	//	"fmt"
	"net/http"

	//"os"
	//	"os/exec"
	"io/ioutil"
	"html/template"
	//	"fmt"
	//	"log"
)

func readHtmlFromFile(fileName string) ([]byte) {

    bs, _ := ioutil.ReadFile(fileName)

    return bs
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(readHtmlFromFile("./index.html"))
}

func DiscosMontados(w http.ResponseWriter ,r *http.Request)  {
	Data:=FormaterDiskInfo(GetInfoSystem())
	t:=template.Must(template.ParseFiles("./discos.html"))
	t.Execute(w,Data)
}

func DiscosDisponibles(w http.ResponseWriter, r *http.Request)  {
	Data:=GetDisks()
	t:=template.Must(template.ParseFiles("./discosDisponibles.html"))
	t.Execute(w,Data)
}


func main() {
	port := ":3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/discos", DiscosMontados)
	mux.HandleFunc("/discosDisponibles", DiscosDisponibles)
	http.ListenAndServe(port, mux)

}
