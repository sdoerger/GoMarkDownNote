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
		fmt.Println(os.Args)
		fmt.Println(os.Args[1])
		fmt.Print("RUNS")

		// FLAGS

		// PRETITLE
		title := ""
		path := "./"
		curDir := ""
		// preTitle := flag.Bool("t", false, "display colorized output")
		// flag.Parse()

		fmt.Println("RUNS")

		// Add filepath as pre title
		path = os.Args[1]
		splitPath := []string(strings.Split(path, "/"))
		curDir = splitPath[len(splitPath)-2]
		noteName := splitPath[len(splitPath)-1]

		title = curDir + ": "

		createdTime := time.Now()

		// TO FIND
		// var specialChar = regexp.MustCompile(`[^a-zA-Z ]`)
		var underscore = regexp.MustCompile(`_`)
		var extension = regexp.MustCompile(`.md`)

		// Replace with
		fileName := noteName
		noteName = underscore.ReplaceAllString(noteName, " ")
		noteName = extension.ReplaceAllString(noteName, "")

		fmt.Println("fileName")
		fmt.Println(fileName)

		// Write MD content
		fileContent := "---" + "\n" + "title: " + "'" + strings.ToUpper(title) + noteName + "'" + "\n" + "created: " + "'" + createdTime.String()[:19] + "'" /*  + "\n"+ "tags: " + "\n" + "\t" + "- " + noteName */ + "\n" + "---" + "\n" + "\n" + "# " + strings.ToUpper(title) + noteName + "\n" + "\n"

		fileTarget := "./" + curDir + "/" + fileName
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
		// 		Params are: File name, Byte Slice item and read/write... righst (0666 is pretty standard)

	}
}
