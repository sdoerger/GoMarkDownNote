package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

func createMarkdown(title string, createdTime time.Time, titlePrefix string) string {
	return fmt.Sprintf(
		"---\n"+
			"title: '%s%s'\n"+
			"created: '%s'\n"+
			"---\n"+
			"\n"+
			"# %s%s\n"+
			"\n",
		strings.ToUpper(titlePrefix),
		title,
		createdTime.String()[:19],
		strings.ToUpper(titlePrefix),
		title,
	)
}

func createCodeBlock(lang string) string {
	return fmt.Sprintf(
		"\n"+
			"```%s\n"+
			"\n"+
			"```"+
			"\n",
		lang,
	)
}

func writeToFile(path, content string) error {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("File exists. Please change file name.\n")
		return nil
	}
	fmt.Println("✅ Note is created.")
	return ioutil.WriteFile(path, []byte(content), 0666)
}

func cleanUpTitle(rawTitle string) string {
	var underscore = regexp.MustCompile(`_`)
	var extension = regexp.MustCompile(`.md`)
	title := underscore.ReplaceAllString(rawTitle, " ")
	title = extension.ReplaceAllString(title, "")
	return title
}

func main() {

	// Define flags
	langPtr := flag.String("lang", "", "language for code block")

	// Parse the flags
	flag.Parse()

	// Get the remaining positional arguments
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No positional arguments provided.")
		return
	}

	path := args[0]
	splitPath := strings.Split(path, "/")
	curDir := splitPath[len(splitPath)-2]
	title := splitPath[len(splitPath)-1]
	titlePrefix := strings.Split(curDir, " ")[0] + ": "
	createdTime := time.Now()

	title = cleanUpTitle(title)

	fileContent := createMarkdown(title, createdTime, titlePrefix)
	if len(*langPtr) > 0 {
		fileContent += createCodeBlock(*langPtr)
	}

	writeToFile(path, fileContent)
}
