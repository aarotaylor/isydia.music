package handlers

import (
	"fmt"
	"net/http"
	"strings"

	ingress "isydia.music/ingress"
	"isydia.music/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Home("Templ App", "...you are now Lost ✨")

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func PurposeHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Home("Templ App", "dfggf ✨")

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

// Function handles the rendering of narrative pages. It will parse the file denoted in the request.
// For now, it reads one file (to be replaced with an objectbox read), and outputs one StoryText block and the episode tagline as a test.
// Eventually, it will render the full narrative with appropriate formatting based on the path after /narratives.
//   - Part of the struct should contain the timestamps in between intake file lines.
func NarrativeHandler(w http.ResponseWriter, r *http.Request) {

	// destination will be a slice of [collection(album), episode], which we can use to determine which narrative file to read and render.
	destination := strings.Split(strings.TrimPrefix(r.URL.Path, "/narratives/"), "/")
	// relative path for the URL will match the relative path of narrative files.
	// urlPath will have /narratives removed.

	// File convention: narrative file types (as well as others)
	// will have spaces replaced with underscores, and will be all lowercase
	collection := strings.ToLower(strings.ReplaceAll(destination[0], " ", "_"))
	item_name := strings.ToLower(strings.ReplaceAll(destination[1], " ", "_"))

	narrative, err := ingress.ParseNarrativeFile("./Narrative/"+collection+"/"+item_name+".txt", 0) // TODO: WRITE NARRATIVE FILE
	if err != nil {
		fmt.Println(err.Error())
		component := views.Home("Dimensional Gateway", "You've found a teapot! ✨")
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}
	}
	// fmt.Printf("[ Parsed narrative: ] %+v\n", narrative)
	component := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, narrative)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}
