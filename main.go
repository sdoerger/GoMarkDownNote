package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	// -----------------
	// Command line Args
	// -----------------

	if os.Args != nil && len(os.Args) > 1 {
		noteName := os.Args[1]
		if noteName != "" {

			// ---
			// title: 'Tailwind: All size image picture Modal/Overlay'
			// updated: 2021-04-14 07:04:24Z
			// created: 2021-04-14 06:59:03Z
			// tags:
			//   - dev.tailwind
			// ---
			// ---

			createdTime := time.Now()

			// TO FIND
			var specialChar = regexp.MustCompile(`[^a-zA-Z ]`)
			var space = regexp.MustCompile(`\s+`)

			// Replace with
			fileName := specialChar.ReplaceAllString(strings.TrimSpace(noteName), "")
			fileName = space.ReplaceAllString(fileName, "_")

			fileContent := "---" + "\n" + "title: " + noteName + "\n" + "updated: " + createdTime.String() + "\n" + "created: " + createdTime.String()[:19] + "\n" + "tags: " + "\n" + "\t" + "- " + noteName + "\n" + "---" + "\n" + "\n" + "# " + noteName + "\n" + "\n"
			fmt.Println(fileContent)

			return
		}
	} else {
		fmt.Println("File Name")
	}

}
