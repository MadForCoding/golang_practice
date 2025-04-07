package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/png", pngHandler)
	http.ListenAndServe(":8899", nil)
}

func pngHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "image/png")
	w.Header().Set("content-disposition", "inline")
	w.WriteHeader(http.StatusOK)
	fmt.Println(os.Getwd())
	http.ServeFile(w, req, "./simple_web/a.png")
}
