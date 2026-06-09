package main

import (
	"log"
	"net/http"

	fsr "isydia.music/handlers"
	ingress "isydia.music/ingress"
	"isydia.music/model"
)

const (
	NarrativeDir = "./Narrative/"
)

func main() {

	third_ring, _ := model.InitBedRock() // initialize objectbox; Syaksa stores a deep history in the third ring.

	// third_ring.
	model.SeedBedRock(NarrativeDir, third_ring) // seed objectbox with Narrative data from the Narrative files. This will be replaced with a more robust seeding process once objectbox is fully implemented.

	mux := http.NewServeMux()
	// ---------------------------
	// cwd, _ := os.Getwd()

	// fmt.Println("working dir:", cwd)
	// if _, err := os.Stat("public/syaksa.css"); err != nil {
	// 	fmt.Println("not found:", err)
	// } else {
	// 	fmt.Println("found syaksa.css")
	// }
	// ---------------------------

	// Routes
	// Static files
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/purpose", fsr.PurposeHandler)

	// Narrative routes
	// these will be of the form
	//   /narratives/[collection(album)]/[episode]
	mux.HandleFunc("/narratives/{collection}", fsr.CollectionHandler)
	mux.HandleFunc("/narratives/{collection}/{episode}", fsr.NarrativeHandler)

	mux.HandleFunc("/updates", fsr.HomeHandler)
	mux.HandleFunc("/releases", fsr.HomeHandler)
	mux.HandleFunc("/hub", fsr.HomeHandler)
	mux.HandleFunc("/credits", fsr.HomeHandler)
	mux.HandleFunc("/", fsr.HomeHandler)

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

	ingress.ParseNarrativeFile("/Narrative/init.txt", 0) // TODO: WRITE NARRATIVE FILE
}
