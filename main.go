package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

		// FLAGS

		// PRETITLE
		titlePrefix := ""
		path := "./"
		curDir := ""
		// preTitle := flag.Bool("t", false, "display colorized output")
		// flag.Parse()

		fmt.Println("RUNS")

		// Add filepath as pre titlePrefix
		path = os.Args[1]
		splitPath := []string(strings.Split(path, "/"))
		curDir = splitPath[len(splitPath)-2]
		title := splitPath[len(splitPath)-1]

		titlePrefix = strings.Split(curDir, " ")[0] + ": "
		createdTime := time.Now()

		// TO FIND
		// var specialChar = regexp.MustCompile(`[^a-zA-Z ]`)
		var underscore = regexp.MustCompile(`_`)
		var extension = regexp.MustCompile(`.md`)

		// Cleanup
		title = underscore.ReplaceAllString(title, " ")
		title = extension.ReplaceAllString(title, "")

		// Write MD content
		fileContent := "---" + "\n" + "title: " + "'" + strings.ToUpper(titlePrefix) + title + "'" + "\n" + "created: " + "'" + createdTime.String()[:19] + "'" /*  + "\n"+ "tags: " + "\n" + "\t" + "- " + title */ + "\n" + "---" + "\n" + "\n" + "# " + strings.ToUpper(titlePrefix) + title + "\n" + "\n"

		fileTarget := path
		if _, err := os.Stat(fileTarget); err == nil {
			fmt.Printf("File exists. Please change file name.\n")
		} else {
			fileContentToByteSlics := []byte(fileContent)
			fmt.Println(fileTarget)
			err := ioutil.WriteFile(fileTarget, fileContentToByteSlics, 0666)
			if err != nil {
				log.Fatalf("Error at writing file: %v", err)
			}
			fmt.Printf("✅: Note is created.\n")
		}

		return
	}
}
