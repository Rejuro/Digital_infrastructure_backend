package controllers

import (
	"fmt"
	"net/http"
)

func Sayhello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет!")
}