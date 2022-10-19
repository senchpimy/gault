package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/sha512"
	"encoding/base64"
)

//var Logged bool

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write(readHtmlFromFile("./login.html"))
	switch r.Method {
	case "GET":
		http.ServeFile(w,r,"./login.html")
	case "POST":
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		fmt.Fprintf(w,"Post form website r.postfrom =%v \n",r.PostForm)
		name:=r.FormValue("name")
		sha_512:=sha512.New()
		sha_512.Write([]byte(name))
		fmt.Fprintf(w,"value = %s\n",name)
		fmt.Fprintf(w,"value in sha_512 = \t%x",base64.StdEncoding.EncodeToString(sha_512.Sum(nil)))

	default: fmt.Fprintf(w,"Error")
	}
}

func main() {
	port := ":3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(port, mux)

}
