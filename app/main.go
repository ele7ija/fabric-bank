package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
    http.HandleFunc("/api/checkhealth", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello, World!")
    })

    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

