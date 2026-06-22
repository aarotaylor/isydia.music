package narrative

import (
	"fmt"

	Shape "isydia.music/model"
)

const (
	NFPATH = "./Narrative/"
)

// This module will read in the Narrative file directory, populating objectbox, and contains
// other functions and methods that will be related to reading narrative files
// toggle the preference for file vs objectbox copy;
func NarrativeIngress() []*Shape.Album {

	var syaksa []*Shape.Narrative
	// ---- these calls directly to the narrative files will be replaced with objectbox reads once objectbox is implemented. ----

	portal, err := ParseNarrativeFile("./Narrative/syaksa/portal.txt", 0) // TODO: WRITE NARRATIVE FILE
	if err != nil {
		fmt.Println("[[ Error reading Narrative file: ]] \n", err.Error())
		// component := views.Home("Dimensional Gateway", "...empty space...")
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

	syaksa = append(syaksa, portal, tff, astraphobia, atomkraft)
	// ---- End Narrative file read into memory ----

	return []*Shape.Album{
		{
			ID:         0,
			AlbumName:  "Syaksa",
			Narratives: syaksa,
			ArtistName: "Isydia",
		},
	}
}

func UltraluminalTrackList() *Shape.Album {

	var album *Shape.Album

	focus_alignment := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Focus Alignment",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 0,
		},
	}

	hadronic_synthesis := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Hadronic Synthesis",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 1,
		},
	}

	owat := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Of Webs and Threads",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 2,
		},
	}

	ov := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Omicron Velorum",
			EpisodeTitle:  "",
			TaglineText:   "",
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
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 5,
		},
	}

	atomkraft := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Atomkraft",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 6,
		},
	}

	ultraluminal := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Ultraluminal",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 7,
		},
	}

	coordinate_space := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Coordinate Space",
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
			TaglineText:   "",
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
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 12,
		},
	}

	ephemera := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Ephemera",
			EpisodeTitle:  "",
			TaglineText:   "",
			Album:         "Ultraluminal",
			EpisodeNumber: 13,
		},
	}

	fsr := &Shape.Narrative{
		Episode: Shape.EpisodeText{
			TrackName:     "Foldspace Router",
			EpisodeTitle:  "",
			TaglineText:   "",
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
