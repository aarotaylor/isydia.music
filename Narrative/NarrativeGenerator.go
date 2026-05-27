package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(r *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := r.ReadString('\n')
	return strings.TrimSpace(text)
}

func readFlags(r *bufio.Reader, name string) []string {
	fmt.Printf("\nDefine up to 4 flags for %s:\n", name)
	flags := []string{}

	for i := 0; i < 4; i++ {
		fmt.Printf("%s flag %d (leave blank to stop): ", name, i+1)
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}
		flags = append(flags, input)
	}

	return flags
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("=== Structured Story Generator ===\n")

	// -------------------------
	// EPISODE METADATA
	// -------------------------
	track := readLine(r, "Track name: ")
	title := readLine(r, "Episode title: ")
	tagline := readLine(r, "Tagline: ")
	album := readLine(r, "Album: ")
	episodeNumber := readLine(r, "Episode number: ")

	// -------------------------
	// FLAGS
	// -------------------------
	aFlags := readFlags(r, "Character")
	bFlags := readFlags(r, "Voice")

	_ = aFlags // reserved for future validation
	_ = bFlags

	// -------------------------
	// STORY LOOP
	// -------------------------
	fmt.Println("\nEnter story entries. Type 'done' to finish.\n")

	storyLines := []string{}
	idCounter := 1

	for {
		fmt.Print("ID (number, 'T', or 'done'): ")
		idInput, _ := r.ReadString('\n')
		idInput = strings.TrimSpace(idInput)

		if idInput == "done" {
			break
		}

		var id string

		if idInput == "T" {
			id = "T"
		} else if idInput == "" {
			id = strconv.Itoa(idCounter)
			idCounter++
		} else {
			id = idInput
		}

		fmt.Print("Character flag: ")
		a, _ := r.ReadString('\n')
		a = strings.TrimSpace(a)

		fmt.Print("Voice flag: ")
		b, _ := r.ReadString('\n')
		b = strings.TrimSpace(b)

		fmt.Print("StoryChunk (text): ")
		c, _ := r.ReadString('\n')
		c = strings.TrimSpace(c)

		line := fmt.Sprintf("%s|%s|%s|%s", id, a, b, c)
		storyLines = append(storyLines, line)

		fmt.Println("Added:", line, "\n")
	}

	// -------------------------
	// OUTPUT FILE
	// -------------------------
	fmt.Print("Output file path: ")
	outPath, _ := r.ReadString('\n')
	outPath = strings.TrimSpace(outPath)

	var builder strings.Builder

	builder.WriteString("# EPISODE\n")
	builder.WriteString(fmt.Sprintf("track: %s\n", track))
	builder.WriteString(fmt.Sprintf("title: %s\n", title))
	builder.WriteString(fmt.Sprintf("tagline: %s\n", tagline))
	builder.WriteString(fmt.Sprintf("album: %s\n", album))
	builder.WriteString(fmt.Sprintf("episode_number: %s\n\n", episodeNumber))

	builder.WriteString("# STORY\n")
	for _, line := range storyLines {
		builder.WriteString(line + "\n")
	}
	builder.WriteString("# END\n")

	err := os.WriteFile(outPath, []byte(builder.String()), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("\nSaved to", outPath)
}
