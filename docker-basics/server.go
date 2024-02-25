package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type responseStatus struct {
	Status string `json:"status"`
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /health request")
	response := responseStatus{Status: "OK"}
	responseStr, _ := json.Marshal(response)
	io.WriteString(w, string(responseStr))
}

func main() {
	http.HandleFunc("/health", handleHealth)

	err := http.ListenAndServe(":8000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
