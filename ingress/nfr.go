package narrative

import (
	"fmt"

	Shape "isydia.music/model"
)

const (
	NFPATH = "./Narrative/"
)

func CollectionSwitch(collection string) *Shape.Album {
	switch collection {
	case "syaksa":
		return GetSyaksa()
	case "scree":
		return GetScree()
	case "makers_i":
		return GetMakers_I()
	case "liminal_spaces":
		return GetLiminalSpaces()
	case "citadel":
		return GetCitadel()
	case "ultraluminal":
		return GetUltraluminal()
	case "midnight_metropolis":
		return GetMM()
	default:
		return nil
	}
}

func StorySwitch(narrative *Shape.Narrative) *Shape.Album {
	switch narrative.Episode.Album {
	case "Portal|Penrose Engine|Astraphobia|Atomkraft|Bulwark|Proton Decay":
		return GetSyaksa()
	case "Scree":
		return GetScree()
	case "Makers I":
		return GetMakers_I()
	case "Liminal Spaces":
		return GetLiminalSpaces()
	case "Citadel":
		return GetCitadel()
	default:
		return nil
	}
}

// This module will read in the Narrative file directory, populating objectbox, and contains
// other functions and methods that will be related to reading narrative files
// toggle the preference for file vs objectbox copy;
func GetSyaksa() *Shape.Album {

	var syaksa []*Shape.Narrative
	album := &Shape.Album{
		ID:         0,
		AlbumName:  "Syaksa",
		ArtistName: "Isydia",
		AlbumText: Shape.EpisodeText{
			TaglineText: "A Cosmic Bulwark",
		},
	}
	// ---- these calls directly to the narrative files will be replaced with objectbox reads once objectbox is implemented. ----

	portal, err := ParseNarrativeFile("./Narrative/syaksa/portal.txt", 0) // TODO: WRITE NARRATIVE FILE
	if err != nil {
		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
		// component := views.Home("Dimensional Gateway", "...empty space...")
	}

	penrose_engine := &Shape.Narrative{
		ID: 2,
		Episode: Shape.EpisodeText{
			TrackName:     "Penrose Engine",
			EpisodeTitle:  "Displacement",
			TaglineText:   "",
			Album:         "Syaksa",
			EpisodeNumber: 2,
		},
	}

	tff, err := ParseNarrativeFile("./Narrative/syaksa/the_forbidden_forest.txt", 0) // TODO: WRITE NARRATIVE FILE
	if err != nil {
		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
		// component := views.Home("Dimensional Gateway", "...empty space...")
	}
	astraphobia, err := ParseNarrativeFile("./Narrative/syaksa/astraphobia.txt", 0) // TODO: WRITE NARRATIVE FILE
	if err != nil {
		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
		// component := views.Home("Dimensional Gateway", "...empty space...")
	}
	atomkraft, err := ParseNarrativeFile("./Narrative/syaksa/atomkraft.txt", 0) // TODO: WRITE NARRATIVE FILE
	if err != nil {
		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
		// component := views.Home("Dimensional Gateway", "...empty space...")
	}

	bulwark := &Shape.Narrative{
		ID: 2,
		Episode: Shape.EpisodeText{
			TrackName:     "Bulwark",
			EpisodeTitle:  "Cycles of Cycles...",
			TaglineText:   "",
			Album:         "Syaksa",
			EpisodeNumber: 6,
		},
	}

	proton_decay := &Shape.Narrative{
		ID: 2,
		Episode: Shape.EpisodeText{
			TrackName:     "Proton Decay",
			EpisodeTitle:  "Oblivion..",
			TaglineText:   "",
			Album:         "Syaksa",
			EpisodeNumber: 7,
		},
	}

	syaksa = append(syaksa, portal, penrose_engine, tff, atomkraft, astraphobia, bulwark, proton_decay)
	// ---- End Narrative file read into memory ----

	album.Narratives = syaksa
	return album
}

func LinksIngress() *Shape.ExternalLinkContainer {

	return nil
}

func GetScree() *Shape.Album {
	var album *Shape.Album

	is := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Induction Shadow",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Scree",
			EpisodeNumber: 0,
		},
	}

	album = &Shape.Album{
		AlbumName:  "Scree",
		ArtistName: "Isydia",
		Narratives: []*Shape.Narrative{is},
	}

	return album
}

func GetMakers_I() *Shape.Album {
	var album *Shape.Album

	manifolds := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Manifold Existence",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Makers I",
			EpisodeNumber: 0,
		},
	}

	flood_basalt := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Flood Basalt",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Makers I",
			EpisodeNumber: 1,
		},
	}

	album = &Shape.Album{
		AlbumName:  "Makers I",
		ArtistName: "Isydia",
		Narratives: []*Shape.Narrative{manifolds, flood_basalt},
	}

	return album
}

func GetLiminalSpaces() *Shape.Album {
	var album *Shape.Album

	mind_the_rift := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Mind the Rift",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Liminal Spaces",
			EpisodeNumber: 0,
		},
	}

	restitution := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Restitution",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Liminal Spaces",
			EpisodeNumber: 1,
		},
	}

	nn := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Nightly Netrunning",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Liminal Spaces",
			EpisodeNumber: 2,
		},
	}

	album = &Shape.Album{
		AlbumName:  "Liminal Spaces",
		ArtistName: "Procyon B",
		Narratives: []*Shape.Narrative{mind_the_rift, restitution, nn},
	}

	return album
}

func GetCitadel() *Shape.Album {
	var album *Shape.Album

	citadel := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Citadel",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Citadel",
			EpisodeNumber: 0,
		},
	}

	outskirts := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Outskirts",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Citadel",
			EpisodeNumber: 1,
		},
	}

	cc := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Corporate Couriers",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Citadel",
			EpisodeNumber: 2,
		},
	}

	album = &Shape.Album{
		AlbumName:  "Citadel",
		ArtistName: "Procyon B",
		Narratives: []*Shape.Narrative{citadel, outskirts, cc},
	}

	return album
}

func GetMM() *Shape.Album {
	var album *Shape.Album

	infiltration := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Infiltration",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Midnight Metropolis",
			EpisodeNumber: 0,
		},
	}

	mm := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Midnight Metropolis",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Midnight Metropolis",
			EpisodeNumber: 1,
		},
	}

	album = &Shape.Album{
		AlbumName:  "Midnight Metropolis",
		ArtistName: "Procyon B",
		Narratives: []*Shape.Narrative{infiltration, mm},
	}

	return album
}

func GetUltraluminal() *Shape.Album {

	var album *Shape.Album

	focus_alignment := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Focus Alignment",
			EpisodeTitle:  "",
			TaglineText:   "Flux Summation",
			Album:         "Ultraluminal",
			EpisodeNumber: 0,
		},
	}

	hadronic_synthesis := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Hadronic Synthesis",
			EpisodeTitle:  "",
			TaglineText:   "E=mc^2 <|> m=E/c^2",
			Album:         "Ultraluminal",
			EpisodeNumber: 1,
		},
	}

	owat := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Of Webs and Threads",
			EpisodeTitle:  "",
			TaglineText:   "Sensory Spinnerets",
			Album:         "Ultraluminal",
			EpisodeNumber: 2,
		},
	}

	ov := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Omicron Velorum",
			EpisodeTitle:  "",
			TaglineText:   "Observers",
			Album:         "Ultraluminal",
			EpisodeNumber: 3,
		},
	}

	plasma_film := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Plasma Film",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 4,
		},
	}

	dipole_moment := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Dipole Moment",
			EpisodeTitle:  "",
			TaglineText:   "Flux Barrier",
			Album:         "Ultraluminal",
			EpisodeNumber: 5,
		},
	}

	atomkraft := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Geochronology",
			EpisodeTitle:  "",
			TaglineText:   "Signal Dissemination",
			Album:         "Ultraluminal",
			EpisodeNumber: 6,
		},
	}

	ultraluminal := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Emergence",
			EpisodeTitle:  "Fractalized Awareness",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 7,
		},
	}

	coordinate_space := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Ariadne's Adytum",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 8,
		},
	}

	soe := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Sea of Exchange",
			EpisodeTitle:  "",
			TaglineText:   "Chaos and Stillness",
			Album:         "Ultraluminal",
			EpisodeNumber: 9,
		},
	}

	relay_interference := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Relay Interference",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 11,
		},
	}

	arafel := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Arafel",
			EpisodeTitle:  "",
			TaglineText:   "Becoming Never Ceases",
			Album:         "Ultraluminal",
			EpisodeNumber: 12,
		},
	}

	ephemera := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Ephemera",
			EpisodeTitle:  "",
			TaglineText:   "Projections and Reflections",
			Album:         "Ultraluminal",
			EpisodeNumber: 13,
		},
	}

	fsr := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Foldspace Router",
			EpisodeTitle:  "",
			TaglineText:   "Windows & Visions",
			Album:         "Ultraluminal",
			EpisodeNumber: 14,
		},
	}

	from_the_noise := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "From the Noise",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 15,
		},
	}

	album = &Shape.Album{

		ID:         3,
		AlbumName:  "Ultraluminal",
		ArtistName: "Procyon B",
		AlbumText: Shape.EpisodeText{
			TrackName:     "",
			EpisodeTitle:  "",
			TaglineText:   "The process of Becoming never ceases...",
			Album:         "Ultraluminal",
			EpisodeNumber: 0,
		},
		Narratives: []*Shape.Narrative{
			focus_alignment,
			hadronic_synthesis,
			owat,
			ov,
			plasma_film,
			dipole_moment,
			atomkraft,
			ultraluminal,
			coordinate_space,
			soe,
			relay_interference,
			arafel,
			ephemera,
			fsr,
			from_the_noise,
		},
	}

	return album
}

// func CollectionIngress() {

// 	var collections []*Shape.Narrative

// 	ultraluminal, err := ParseNarrativeFile("./Narrative/ultraluminal/ultraluminal_album.txt", 0) // TODO: WRITE NARRATIVE FILE
// 	if err != nil {
// 		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
// 		// component := views.Home("Dimensional Gateway", "...empty space...")
// 	}

// 	syaksa, err := ParseNarrativeFile("./Narrative/ultraluminal/ultraluminal_album.txt", 0) // TODO: WRITE NARRATIVE FILE
// 	if err != nil {
// 		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
// 		// component := views.Home("Dimensional Gateway", "...empty space...")
// 	}
// }
