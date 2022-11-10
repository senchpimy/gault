package main

import (
    "fmt"
    "net/http"
"encoding/json"
)

type InputFromPage struct {
    Input string
}

func HandleButtons(w http.ResponseWriter, r *http.Request){
fmt.Println(r.Form)
 decoder := json.NewDecoder(r.Body)
    var t InputFromPage
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    fmt.Println(t.Input)
    switch t.Input{
	case "logout":
		http.Redirect(w, r, "https://www.google.com", http.StatusFound)
		logoutHandler(w,r)
    }
}


