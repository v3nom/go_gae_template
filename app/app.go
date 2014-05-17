package main

import (
	"fmt"
	"github.com/v3nom/go_gae_template/service"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, service.GetGreeting())
}
