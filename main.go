package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var matchingRegexp, replaceString string
	//dir := "dir/to/walk"
	dir := "."

	switch len(argsWithoutProg) {
	case 2:
		matchingRegexp, replaceString = argsWithoutProg[0], argsWithoutProg[1]
	case 3:
		dir, matchingRegexp, replaceString = argsWithoutProg[0], argsWithoutProg[1], argsWithoutProg[2]
	default:
		fmt.Println("usage: main [<path>] \"<match_regexp>\" \"<replace_string>\"")
		os.Exit(1)
	}

	//matchingRegexp = "^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$"
	//replaceString = "$2 - $1 - $3 of $4.$5"
	re, err := regexp.Compile(matchingRegexp)
	if err != nil {
		fmt.Println("There was an error with the provided regular expresion. See below for more info.")
		fmt.Printf("\n%v\n\n", err)
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		if info.IsDir() {
			return nil
		} else if re.MatchString(info.Name()) {
			dir := filepath.Dir(path)
			newName := re.ReplaceAllString(info.Name(), replaceString)
			newPath := filepath.Join(dir, newName)
			// Chek for conflicts between oldPath and newPath
			if _, err := os.Stat(newPath); !os.IsNotExist(err) {
				fmt.Printf("Skipping file \"%s\". A file with the proposed name \"%s\" already exists.\n", path, newPath)
				return nil
			}
			// If not, rename the file
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Println("There was an error renaming the file. See below for more info.")
				fmt.Printf("\n%v\n\n", err)
			}
			fmt.Printf("File \"%s\" renamed to \"%s\" correctly.\n", path, newPath)
			//fmt.Printf("File \"%s\" in \"%s\" matches the provided regexp\n", info.Name(), dir)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}
