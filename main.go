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
		fmt.Println(len(os.Args))
		fmt.Println(os.Args[len(os.Args)-1])
		noteName := os.Args[len(os.Args)-1]
		if noteName != "" {

			// FLAGS

			// PRETITLE
			title := ""
			path := "./"
			curDir := ""
			// preTitle := flag.Bool("t", false, "display colorized output")
			// flag.Parse()

			if len(os.Args) > 2 {

				// Add filepath as pre title
				path = os.Args[1]
				splitPath := []string(strings.Split(path, "/"))
				curDir = splitPath[len(splitPath)-1]
				fmt.Println("curDir")
				fmt.Println(curDir)
				title = curDir + ": "
			}
			fmt.Println(title)

			createdTime := time.Now()

			// TO FIND
			var specialChar = regexp.MustCompile(`[^a-zA-Z ]`)
			var space = regexp.MustCompile(`\s+`)

			// Replace with
			noteTitle := specialChar.ReplaceAllString(strings.TrimSpace(noteName), "")
			fileName := space.ReplaceAllString(noteTitle, "_")

			// Write MD content
			fileContent := "---" + "\n" + "title: " + "'" + strings.ToUpper(title) + noteName + "'" + "\n" + "updated: " + "'" + createdTime.String() + "'" + "\n" + "created: " + "'" + createdTime.String()[:19] + "'" /*  + "\n"+ "tags: " + "\n" + "\t" + "- " + noteName */ + "\n" + "---" + "\n" + "\n" + "# " + strings.ToUpper(title) + noteName + "\n" + "\n"

			fileTarget := "./" + curDir + "/" + fileName + ".md"
			if _, err := os.Stat(fileTarget); err == nil {
				fmt.Printf("File exists. Please change file name.\n")
			} else {
				fileContentToByteSlics := []byte(fileContent)
				fmt.Println(fileTarget)
				err := ioutil.WriteFile("./"+curDir+"/"+fileName+".md", fileContentToByteSlics, 0666)
				if err != nil {
					log.Fatalf("Error at writing file: %v", err)
				}
				fmt.Printf("✅: Note is created.\n")
			}

			return
			// 		Params are: File name, Byte Slice item and read/write... righst (0666 is pretty standard)
		}
	} else {
		fmt.Println("File Name")
	}

}
