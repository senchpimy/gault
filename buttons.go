package main

import (
    "fmt"
    "net/http"
"encoding/json"
)

type InputFromPage struct {
    Input string
}

type Response struct {
    Output string
}

func HandleButtons(w http.ResponseWriter, r *http.Request){
    decoder := json.NewDecoder(r.Body)
    var t InputFromPage
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    fmt.Println(t.Input)
    switch t.Input{
	case "logout":
		logoutHandler(w,r)
		loggedOut:= Response{Output:"Done"}
		byteArray, err := json.Marshal(loggedOut)
        	if err != nil {fmt.Println(err)}
		w.Write(byteArray)
	default:
		fmt.Println("Input no encontrado")
    }
}


