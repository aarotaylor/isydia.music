package main

import (
	"log"
	"net/http"

	fsr "isydia.music/handlers"
	ingress "isydia.music/ingress"
)

const (
	NarrativeDir = "./Narrative/"
)

func main() {

	// third_ring, _ := model.InitBedRock() // initialize objectbox; Syaksa stores a deep history in the third ring.

	// // third_ring.
	// model.SeedBedRock(NarrativeDir, third_ring) // seed objectbox with Narrative data from the Narrative files. This will be replaced with a more robust seeding process once objectbox is fully implemented.

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

	mux.HandleFunc("/updates", fsr.RiftHandler)
	mux.HandleFunc("/releases", fsr.RiftHandler)
	mux.HandleFunc("/hub", fsr.HubHandler)
	mux.HandleFunc("/credits", fsr.AttributeHandler)

	mux.HandleFunc("/isydia", fsr.IsydiaHandler)
	mux.HandleFunc("/procyon_b", fsr.ProcyonHandler)
	// mux.HandleFunc("/ingress", fsr.NexusHandler)
	mux.HandleFunc("/", fsr.NexusHandler)

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

	ingress.ParseNarrativeFile("/Narrative/init.txt", 0) // TODO: WRITE NARRATIVE FILE
}
