package main

import (
	"net/http"
	"io/ioutil"
	"html/template"
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

func SambaConfiguration(w http.ResponseWriter, r *http.Request)  {
	Configuration:=GetAllConfigurations()
	t:=template.Must(template.ParseFiles("./samba.html"))
	t.Execute(w,Configuration)
}

func main() {
	port := ":3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/discos", DiscosMontados)
	mux.HandleFunc("/discosDisponibles", DiscosDisponibles)
	mux.HandleFunc("/SambaConfi", SambaConfiguration)
	http.ListenAndServe(port, mux)

}
