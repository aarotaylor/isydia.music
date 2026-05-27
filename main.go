package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	fsr "isydia.music/handlers"
	ingress "isydia.music/ingress"
)

func main() {

	mux := http.NewServeMux()
	cwd, _ := os.Getwd()
	fmt.Println("working dir:", cwd)
	if _, err := os.Stat("public/syaksa.css"); err != nil {
		fmt.Println("not found:", err)
	} else {
		fmt.Println("found syaksa.css")
	}
	// Routes
	// Static files
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/purpose", fsr.PurposeHandler)

	// Narrative routes
	// these will be of the form
	//   /narratives/[collection(album)]/[episode]
	mux.HandleFunc("/narratives/", fsr.NarrativeHandler)
	mux.HandleFunc("/", fsr.HomeHandler)

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

	ingress.ParseNarrativeFile("/Narrative/init.txt", 0) // TODO: WRITE NARRATIVE FILE
}
