package model

import "fmt"

// One StoryText block is roughly a paragraph, sometimes more, sometimes less.
// A slice of StoryText objects form a Narrative.
type StoryText struct {
	ID      uint64 `json:"id"`
	Speaker string // Naomi, Kai, Emissary, etc.
	Voice   string // Spoken, Thought
	Text    string // "It came on us at once, like dropping weight on our backs..."
	Episode int    // Order of occurrence within the page
}

// Object that contains a timestamp, along with a slice of StoryText objects that occur at that timestamp.
// Timestamp given as an int, to be converted to mm:ss in the frontend. For example, 90000 that would be 1:30.
type Anchor struct {
	ID        uint64 `json:"id"`
	Timestamp int    // remember to convert to multiple of 1000 when writing to narrative file
	Scene     string // description of this moment in the episode, e.g. "Theory of Mind"
	Sequence  []StoryText
}

// All text for a given page.
// Narrative.String() should return the full text of the Narrative, with appropriate formatting.
type Narrative struct {
	ID         uint64   `json:"id"`
	Collection []Anchor // Slice of StoryText blocks, ordered by timestamp
	Episode    EpisodeText
}

// Header text for a given episode. Placed at the top of a Narrative
type EpisodeText struct {
	ID            uint64 `json:"id"`
	TrackName     string // Name of the track associated with this episode
	EpisodeTitle  string // Title of the episode
	TaglineText   string
	Album         string
	EpisodeNumber int
}

// General object for page text that isn't Narrative related
// e.g. text on the Purpose page
type PageText struct {
	ID   uint64 `json:"id"`
	Text string
}

func SecondsToMMSS(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60

	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}

type Album struct {
	ID         uint64 `json:"id"`
	AlbumName  string
	Narratives []Narrative
	ArtistName string
}
