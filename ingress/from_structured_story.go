package narrative

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	Shape "isydia.music/model"
)

func ParseNarrativeFile(path string, id uint64) (*Shape.Narrative, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n := &Shape.Narrative{
		ID: id,
	}

	var inStory bool
	var story []Shape.Anchor
	var sequence []Shape.StoryText
	// var anchors []Shape.Anchor
	// var anchor Shape.Anchor
	var episode Shape.EpisodeText
	var timestamp int
	var scene string

	//scannerLoop:
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// fmt.Println("[ RAW ] ", line)

		switch {
		case strings.HasPrefix(line, "# EPISODE"):
			inStory = false

		case strings.HasPrefix(line, "# STORY"):
			inStory = true

		case strings.HasPrefix(line, "# END"):
			// on the last line of the narrative file, we need to append the final anchor to the story before exiting the loop, since there won't be another timestamp line to trigger the append.
			inStory = false
			story = append(story, Shape.Anchor{
				Timestamp: timestamp,
				Scene:     scene,
				Sequence:  sequence,
			})

		default:

			if !inStory {
				// episode metadata
				if after, ok := strings.CutPrefix(line, "track:"); ok {
					episode.TrackName = strings.TrimSpace(after)
				}
				if after, ok := strings.CutPrefix(line, "title:"); ok {
					episode.EpisodeTitle = strings.TrimSpace(after)
				}
				if after, ok := strings.CutPrefix(line, "tagline:"); ok {
					episode.TaglineText = strings.TrimSpace(after)
				}
			} else {
				// story line
				parts := strings.SplitN(line, "|", 4)
				if len(parts) < 4 {
					return nil, errors.New("invalid story line: " + line)
				}

				// this block currently associates the timestamp *after* the StoryText lines, but it should be associated *before* the StoryText lines.
				// This is because the timestamp indicates when the StoryText lines occur in the episode, so it should be read first before reading the StoryText lines that occur at that timestamp.
				if parts[0] == "T " {

					// read the timestamp and assign it to the current Anchor.
					// make the current anchor the destination for new StoryText objects until the next timestamp is read.

					// on the first timestamp, there won't be a current anchor to assign the StoryText objects to, so we can skip the assignment step and just read the timestamp and create the first anchor.

					// -------------
					if parts[1] != "0" {
						story = append(story, Shape.Anchor{
							Timestamp: timestamp,
							Scene:     scene,
							Sequence:  sequence,
						})
						sequence = nil // reset story for the next anchor
						scene = ""
					} else {
						// if timestamp is 0, it means the StoryText lines that follow occur before the first timestamp, so we can just read the StoryText lines without creating an anchor until we read the first timestamp.
						// This is because the timestamp indicates when the StoryText lines occur in the episode, so if we read StoryText lines before reading the first timestamp, we won't know when they occur in the episode and won't be able to associate them with an anchor.
					}
					timestamp, err = strconv.Atoi(strings.TrimSpace(parts[1]))
					if err != nil {
						return nil, err
					}
					scene = strings.TrimSpace(parts[2])

					// -------------
				} else {
					epNum, err := strconv.Atoi(strings.TrimSpace(parts[0]))
					if err != nil {
						return nil, err
					}

					sequence = append(sequence, Shape.StoryText{
						Speaker: strings.TrimSpace(parts[1]),
						Voice:   strings.TrimSpace(parts[2]),
						Text:    strings.TrimSpace(parts[3]),
						Episode: epNum,
					})
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	n.Collection = story
	n.Episode = episode

	return n, nil
}
