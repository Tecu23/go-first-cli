package ls

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

const colorReset = "\033[0m"
const colorBlue = "\033[0;31m"

func Ls() {
	show_hidden_files := flag.Bool("a", false, "Display hidden files")
	flag.Parse()

	multiple_directories := false

	args := flag.Args()

	if len(args) > 1 {
		multiple_directories = true
	}

	for _, s := range args {

		files, err := getFilesFromFolder(s, *show_hidden_files)

		if err != nil {
			fmt.Println("ls:", err)
			continue
		}

		if multiple_directories {
			fmt.Printf("%s\n", s)
		}

		for _, f := range files {
			fmt.Printf("%s ", f)
		}
		fmt.Println()
	}
}

func getFilesFromFolder(path string, show_hidden_files bool) ([]string, error) {
	fileNames := make([]string, 0)

	if show_hidden_files {
		fileNames = append(fileNames, ".")
		fileNames = append(fileNames, "..")
	}

	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !show_hidden_files && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fileNames = append(fileNames, file.Name())
	}

	sort.Strings(fileNames)

	return fileNames, nil
}
