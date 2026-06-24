package handlers

import (
	"fmt"
	"net/http"
	"strings"

	ingress "isydia.music/ingress"
	narrative "isydia.music/ingress"
	"isydia.music/model"
	"isydia.music/views"
)

// Needs a helper function that initializes the narrative texts present in Narrative/ into memory.
// And/or objectbox

func RiftHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Rift("⌞⁂⌝", "...you are now Lost ✨")

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}

}

func NexusHandler(w http.ResponseWriter, r *http.Request) {
	component := views.AspectDisplay("「⌞---- ⁂ ----⌝」", "...pathways converge here...")

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func PurposeHandler(w http.ResponseWriter, r *http.Request) {

	component := views.Home("Dimensional Gateway", "...empty space...")
	purposeText, err := ingress.ParseNarrativeFile("./Narrative/purpose.txt", 0) // TODO: WRITE NARRATIVE FILE
	if err != nil {
		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}
		return
	}

	component = views.BasicText(purposeText.Collection[0].Sequence[0].Text) // for now, just render the first StoryText block of the first Anchor in the Narrative. To be replaced with a more robust rendering of the full Narrative.
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func IsydiaHandler(w http.ResponseWriter, r *http.Request) {

	component := views.Home("__", "...")

	syaksa := narrative.GetSyaksa()
	makers := narrative.GetMakers_I()
	scree := narrative.GetScree()

	discog := []*model.Album{syaksa, makers, scree}

	component = views.IsydiaLayout(discog) // for now, just render the first StoryText block of the first Anchor in the Narrative. To be replaced with a more robust rendering of the full Narrative.
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func ProcyonHandler(w http.ResponseWriter, r *http.Request) {

	component := views.Home("__", "...")
	// var discog []*model.Album{}

	ultraluminal := narrative.UltraluminalTrackList()
	mm := narrative.GetMM()
	liminal_spaces := narrative.GetLiminalSpaces()

	discog := []*model.Album{ultraluminal, mm, liminal_spaces}

	component = views.ProcyonLayout(discog) // for now, just render the first StoryText block of the first Anchor in the Narrative. To be replaced with a more robust rendering of the full Narrative.
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func CollectionHandler(w http.ResponseWriter, r *http.Request) {

	// collection, item_name := pathToNames(r.URL.Path)
	collection := r.PathValue("collection")

	syaksa := ingress.GetSyaksa()

	makers := ingress.GetMakers_I()
	citadel := ingress.GetCitadel()
	mm := ingress.GetMM()
	scree := ingress.GetScree()
	liminal_spaces := ingress.GetLiminalSpaces()
	// ---- End Narrative file read into memory ----
	ultraluminal := ingress.UltraluminalTrackList()

	// cmp := views.CollectionLayout(syaksa)

	//-------------------------------------------------
	// narrative, err := ingress.ParseNarrativeFile("./Narrative/collections/"+collection+".txt", 0) // TODO: WRITE NARRATIVE FILE
	// if err != nil {
	// 	fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())

	// 	cmp := views.Home("Dimensional Gateway", "...empty space...")
	// 	if err := cmp.Render(r.Context(), w); err != nil {
	// 		http.Error(w, "Render error", http.StatusInternalServerError)
	// 	}
	// }

	// in switch statements, inside the switch block, the proper syntax calls for case statements to be at the same indentation level as the switch statement, and the code inside each case to be indented one level further. The break statement is not needed in Go, as it automatically breaks after each case unless you explicitly use fallthrough.
	// the default statement is used to handle any cases that are not explicitly handled by the case statements. It should be placed at the end of the switch block, and it will execute if none of the case statements match the switch expression.

	switch collection {
	case "syaksa":
		fmt.Println("Isydia - [ " + collection + " ]")
		// cmp := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Isydia", narrative)
		cmp := views.CollectionLayout(syaksa)
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}
		return
	case "makers_i":
		fmt.Println("Isydia - [ " + collection + " ]")
		// cmp := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Isydia", narrative)
		cmp := views.CollectionLayout(makers)
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}
		return
	case "scree":
		fmt.Println("Isydia - [ " + collection + " ]")
		// cmp := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Isydia", narrative)
		cmp := views.CollectionLayout(scree)
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}
		return
	case "liminal_spaces":

		fmt.Println("Procyon B - [ " + collection + " ]")
		// cmp := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Procyon B", narrative)
		cmp := views.CollectionLayout(liminal_spaces)
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)

		}
		return
	case "ultraluminal":

		fmt.Println("Procyon B - [ " + collection + " ]")
		// cmp := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Procyon B", narrative)
		cmp := views.CollectionLayout(ultraluminal)
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)

		}
		return
	case "citadel":

		fmt.Println("Procyon B - [ " + collection + " ]")
		// cmp := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Procyon B", narrative)
		cmp := views.CollectionLayout(citadel)
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)

		}
		return
	case "midnight_metropolis":

		fmt.Println("Procyon B - [ " + collection + " ]")
		// cmp := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Procyon B", narrative)
		cmp := views.CollectionLayout(mm)
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)

		}
		return
	default:

		fmt.Println("Unknown collection: [ " + collection + " ]")
		// This type isn't implemented yet. It will wait until objectbox is implemented.
		// Objectbox will store Albums, which contain Narratives.
		cmp := views.Home("Dimensional Gateway", "...you have encountered the Absence of something...")
		if err := cmp.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}
	}
}

func pathToNames(path string) (string, string) {
	// destination will be a slice of [collection(album), episode], which we can use to determine which narrative file to read and render.
	destination := strings.Split(strings.TrimPrefix(path, "/narratives/"), "/")
	// relative path for the URL will match the relative path of narrative files.
	// urlPath will have /narratives removed.

	// File convention: narrative file types (as well as others)
	// will have spaces replaced with underscores, and will be all lowercase

	collection := strings.ToLower(strings.ReplaceAll(destination[0], " ", "_"))
	item_name := strings.ToLower(strings.ReplaceAll(destination[1], " ", "_"))

	return collection, item_name
}

// Function handles the rendering of narrative pages. It will parse the file denoted in the request.
// For now, it reads one file (to be replaced with an objectbox read), and outputs one StoryText block and the episode tagline as a test.
// Eventually, it will render the full narrative with appropriate formatting based on the path after /narratives.
//   - Part of the struct should contain the timestamps in between intake file lines.
func NarrativeHandler(w http.ResponseWriter, r *http.Request) {

	// the declaration of narrative and err here allows them to be used in the if statements below, while also allowing them to be assigned values within those if statements. This is a common pattern in Go for handling errors that may occur during the execution of a function, while still allowing the variables to be accessible outside of the if statement blocks.
	//
	var narrative *model.Narrative
	var err error

	collection := r.PathValue("collection")
	item_name := r.PathValue("episode")
	fmt.Println(item_name)
	// Handle the case where only the collection is selected, it should list the episodes in that collection.
	if item_name == "" {

		narrative, err = ingress.ParseNarrativeFile("./Narrative/"+collection+".txt", 0) // TODO: WRITE NARRATIVE FILE
		if err != nil {
			fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())

			component := views.Home("Dimensional Gateway", "...empty space...")
			if err := component.Render(r.Context(), w); err != nil {
				http.Error(w, "Render error", http.StatusInternalServerError)
			}
		}

	} else {

		narrative, err = ingress.ParseNarrativeFile("./Narrative/"+collection+"/"+item_name+".txt", 0) // TODO: WRITE NARRATIVE FILE
		if err != nil {
			fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())

			cmp := views.Home("Dimensional Gateway", "...empty space...")
			if err := cmp.Render(r.Context(), w); err != nil {
				http.Error(w, "Render error", http.StatusInternalServerError)
			}
		}

	}

	switch collection {
	case "syaksa", "manifold_existence":
		component := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Isydia", narrative)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}
	case "liminal_spaces", "ultraluminal":
		component := views.NarrativePage(narrative.Episode.EpisodeTitle, item_name, "Procyon B", narrative)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)

		}
	default:
		// This type isn't implemented yet. It will wait until objectbox is implemented.
		// Objectbox will store Albums, which contain Narratives.
		component := views.Home("Dimensional Gateway", "...you have encountered the Absence of something... ✨")
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Render error", http.StatusInternalServerError)
		}

	}
	// fmt.Printf("[ Parsed narrative: ] %+v\n", narrative)
}
