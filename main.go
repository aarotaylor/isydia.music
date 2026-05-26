package main

import (
	"log"
	"net/http"

	fsr "isydia.music/handlers"
	ingress "isydia.music/ingress"
)

func main() {

	mux := http.NewServeMux()

	// Routes
	// Static files
	fs := http.FileServer(http.Dir("public"))
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
