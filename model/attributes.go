package model

// based on the narrative, return all assets associated with it.
// initially will be a switch, but eventually will pull from Objectbox
func (n *Album) GetAssets() []*Asset {

	switch n.AlbumName {
	case "Ultraluminal":
		return []*Asset{
			{
				Creator:   "VimesArt",
				License:   "CC BY-NC 4.0",
				Filename:  "ultraluminal_cover.png",
				Notes:     "Cover art for Ultraluminal",
				Type:      AssetTypeImage,
				SourceURL: "https://vimesart.com/projects/XJybaR",
			},
			{
				Creator:   "Liu Zishan",
				License:   "CC BY-NC 4.0",
				Filename:  "/static/img/relay_interference.jpg",
				Notes:     "Cover art for Relay Interference",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
			{
				Creator:   "Liu Zishan",
				License:   "CC BY-NC 4.0",
				Filename:  "/static/img/omicron_velorum [ liu-zishan ].jpg",
				Notes:     "Cover art for Omicron Velorum",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
		}
	case "Syaksa":
		return []*Asset{
			{
				Creator:   "Tithi Luadthong",
				License:   "CC BY-NC 4.0",
				Filename:  "static/img/SYAKSA.jpg",
				Notes:     "Cover art for Syaksa",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
			{
				Creator:   "GrandeDuc",
				License:   "CC BY-NC 4.0",
				Filename:  "static/img/subterrine-square_grandeduc.jpg",
				Notes:     "Cover art for Penrose Engine",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
			{
				Creator:   "Tithi Luadthong",
				License:   "CC BY-NC 4.0",
				Filename:  "/static/img/_portal.jpg",
				Notes:     "Cover art for Portal",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
			{
				Creator:   "Tithi Luadthong",
				License:   "CC BY-NC 4.0",
				Filename:  "/static/img/_atomkraft.jpg",
				Notes:     "Cover art for Atomkraft",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},

			{
				Creator:   "Tithi Luadthong",
				License:   "CC BY-NC 4.0",
				Filename:  "/static/img/_astraphobia.jpg",
				Notes:     "Cover art for Astraphobia",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
			{
				Creator:   "Tithi Luadthong",
				License:   "CC BY-NC 4.0",
				Filename:  "/static/img/_proton-decay.jpg",
				Notes:     "Cover art for Proton Decay",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
			{
				Creator:   "Tithi Luadthong",
				License:   "CC BY-NC 4.0",
				Filename:  "/static/img/_Across the Shifting Sea.jpg",
				Notes:     "Cover art for Across the Shifting Sea",
				Type:      AssetTypeImage,
				SourceURL: "@grandfailure9",
			},
			{
				Creator:   "VimesArt",
				License:   "CC BY-NC 4.0",
				Filename:  "static/img/manifold_existence-[vimesart].png",
				Notes:     "Cover art for Manifold Existence",
				Type:      AssetTypeImage,
				SourceURL: "https://vimesart.com/projects/AZykmz",
			},
		}
	}

	return []*Asset{}
}
