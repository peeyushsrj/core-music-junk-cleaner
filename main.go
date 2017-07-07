/*Program to remove junk data in mp3 files*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	snip "github.com/peeyushsrj/golang-snippets"
)

var (
	webRex         = "(www.|)[a-zA-Z0-9_\\-]+\\.[a-zA-Z]{2,4}"
	fileType       = ".mp3"
	shouldRename   = false
	junkFilesFound = false
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: [path to mp3s] [--rename]\n")
		return
	} else if len(os.Args) == 3 {
		//Checking for rename flag
		if os.Args[2:][0] == "--rename" {
			shouldRename = true
		}
	}

	//Read Junk List stored in array
	junkList, err := snip.ReadLineFromFile("./junk.txt")
	if err != nil {
		log.Fatal("Error loading in junklist", err)
	}
	//Read Music List stored in array
	musicList, err := snip.BrowseXFiles(fileType, os.Args[1:][0])
	if err != nil {
		log.Fatal("Error in browsing mp3 files", err)
	}
	if len(musicList) == 0 {
		fmt.Println("No Music Files here!...")
		return
	}
	//Regex for website junk
	rx, _ := regexp.Compile(webRex)

	for _, fi := range musicList {
		//This contains base file name i.e. without directory name
		//also removed .mp3 extension
		fiCleaned := filepath.Base(strings.TrimSuffix(fi, fileType))
		//Potential junk by above regex
		if rx.MatchString(fiCleaned) {
			//Flag for existance of junk files
			if !junkFilesFound {
				junkFilesFound = true
			}
			junk := snip.StringInSlice(fiCleaned, junkList)
			//If Junk not found in the list
			if junk == "" {
				fmt.Println("Enter the spam part for : ", fiCleaned)
				//Empty response or space response handled by scanf
				_, err := fmt.Scanln(&junk)
				if err != nil {
					log.Fatal("Error in reading junk variable: ", err)
				}
				junkList = append(junkList, junk)

				err = snip.AppendStringToFile(junk, "junk.txt", true)
				if err != nil {
					log.Fatal("Error appending to junk.txt", err)
				}
			}
			//Here we go junk variable set & added to memory
			if shouldRename {
				errt := os.Rename(fi, strings.Replace(fi, junk, "", 1))
				if errt != nil {
					log.Println("Error in re-naming: ", errt)
					continue
				}
			}
			fmt.Println("Old Name: ", fi)
			fmt.Println("New Name: ", strings.Replace(fi, junk, "", 1))
		}
	}
	if !junkFilesFound {
		fmt.Println("No Junk Files Found!")
	} else {
		fmt.Println("Process Completed Successfully!")
	}
}
