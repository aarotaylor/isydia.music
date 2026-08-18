package model

import (
	"fmt"
	"os"
	"path/filepath"
)

// ObjectBox initialization
// func InitBedRock() (*objectbox.ObjectBox, error) {
// 	objectBox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()
// 	if err != nil {
// 		fmt.Println("Error initializing objectbox: ", err.Error())
// 		return nil, err
// 	}
// 	return objectBox, nil
// }

// Seed objectbox with the Narrative Files. Skips seeding if the Narrative data already exists in objectbox, to avoid duplicates. This will be replaced with a more robust seeding process once objectbox is fully implemented.
func SeedBedRock(narratives string /*, objectBox *objectbox.ObjectBox */) {
	// For each Narrative file, check if the corresponding Narrative data already exists in objectbox. If not, parse the Narrative file and write the data to objectbox.

	filepath.Walk(narratives, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking through narrative directory: ", err.Error())
			return err
		}
		fmt.Println(path, info.Name())
		return nil
	})
}

func GetIsydiaLinks() []*ExternalLink {
	links := []*ExternalLink{
		{ID: 0,
			Name:        "Bandcamp",
			URL:         "https://isydia.bandcamp.com/",
			Description: "Listen to and purchase Isydia's music on Bandcamp.",
		},
		{ID: 2,
			Name:        "Soundcloud",
			URL:         "https://soundcloud.com/isydia",
			Description: "Listen to Isydia's music on Soundcloud.",
		}, {ID: 3,
			Name:        "Spotify",
			URL:         "https://open.spotify.com/artist/3Zq3XlTsFGZEWeYnReIofr?si=e-axdGRsS_aOkRFO52_I3Q",
			Description: "Listen to Isydia's music on Spotify.",
		}, {ID: 4,
			Name:        "Apple Music",
			URL:         "https://music.apple.com/us/artist/isydia/1541548981",
			Description: "Listen to Isydia's music on Apple Music.",
		}, {ID: 1,
			Name:        "SpaceAmbient",
			URL:         "https://youtube.com/playlist?list=PLuqZTQ_5gspoSUpWCqcpqKns7EFPH2pYN&si=GWS1qzgwfyd6IOrj",
			Description: "Listen to Isydia's music on the SpaceAmbient YouTube channel.",
		}}

	return links
}

func GetProcyonLinks() []*ExternalLink {
	links := []*ExternalLink{
		{ID: 0,
			Name:        "Bandcamp",
			URL:         "https://isydia.bandcamp.com/",
			Description: "Listen to and purchase Procyon B's music on Bandcamp.",
		},
		{ID: 2,
			Name:        "Soundcloud",
			URL:         "https://soundcloud.com/procyon_b",
			Description: "Listen to Procyon's music on Soundcloud.",
		}, {ID: 3,
			Name:        "Spotify",
			URL:         "https://open.spotify.com/artist/3Zq3XlTsFGZEWeYnReIofr?si=e-axdGRsS_aOkRFO52_I3Q",
			Description: "Listen to Procyon's music on Spotify.",
		}, {ID: 4,
			Name:        "Apple Music",
			URL:         "https://music.apple.com/us/artist/procyon-b/1510628527",
			Description: "Listen to Procyon B's music on Apple Music.",
		}, {ID: 1,
			Name:        "SpaceAmbient",
			URL:         "https://youtube.com/playlist?list=PLuqZTQ_5gspoSUpWCqcpqKns7EFPH2pYN&si=GWS1qzgwfyd6IOrj",
			Description: "Listen to Procyon B's music on the SpaceAmbient YouTube channel.",
		}}

	return links
}

func GetSpaceAmbientLinks() []*ExternalLink {

	links := []*ExternalLink{
		{ID: 0,
			Name:        "Flood Basalt",
			URL:         "https://music.youtube.com/watch?v=yH3swz1vBeg",
			Description: "Listen to Flood Basalt on the SpaceAmbient channel.",
		}, {ID: 1,
			Name:        "Syaksa",
			URL:         "https://www.youtube.com/watch?v=PoCv2IA3Tz0",
			Description: "Listen to Syaksa on the SpaceAmbient channel.",
		},
	}
	return links
}

// func GetObjectBox() (*objectbox.ObjectBox, error) {

// }
//
