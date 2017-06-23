package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ResponseMsg struct {
	Count   int
	Context string
}

var (
	musicList []string
	spamList  []string
	upgrader  = websocket.Upgrader{}
)

func LinesFromFile(path string) ([]string, error) {
	var arr []string

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, strings.TrimSpace(scanner.Text()))
	}
	if scanner.Err() != nil {
		return arr, scanner.Err()
	}
	return arr, nil
}

func BrowseXFiles(x string, root string) ([]string, error) {
	var arr []string
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			if strings.HasSuffix(f.Name(), x) { //.mp3
				arr = append(arr, path)
			}
		}
		return nil
	})
	if err != nil {
		return arr, err
	}
	return arr, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [path to mp3s]\n", os.Args[0])
		return
	}

	//load local spam list
	l, err := LinesFromFile("./spam.txt")
	if err != nil {
		log.Fatal("Error loading in spamlist", err)
	}
	spamList = l

	// musicList = []string{}
	musicList, err = BrowseXFiles(".mp3", os.Args[1:][0])
	if err != nil {
		log.Fatal("Error in walking over files", err)
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/ws", compute)
	log.Println("Running on :7899")
	err = http.ListenAndServe(":7899", nil)
	if err != nil {
		log.Fatal("listenAndServe", err)
	}
}

func home(rw http.ResponseWriter, req *http.Request) {
	// fmt.Println("Client connected", req.RemoteAddr)
	var v = struct {
		Host  string
		Count int
	}{
		req.Host,
		len(musicList),
	}
	t := template.Must(template.ParseFiles("socketed.html"))
	t.Execute(rw, &v)
}

func compute(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	// defer c.Close()

	//Regex for websites spam
	webRex := "(www.|)[a-zA-Z0-9_\\-]+\\.[a-zA-Z]{2,4}"
	rx, _ := regexp.Compile(webRex)
	var iter = 0
	var cleaned_fi string

	//Scan music list
	for _, fi := range musicList {
		iter = iter + 1

		//Only base names, mp3 extension exclude
		cleaned_fi = filepath.Base(strings.TrimSuffix(fi, ".mp3"))

		//Possible spam with websites name, other exclude
		if rx.MatchString(cleaned_fi) {

			//Spam List Match
			spam := stringInSlice(cleaned_fi, spamList)
			if spam == "" { //spam not found
				c.WriteJSON(&ResponseMsg{iter, cleaned_fi})

				var v = struct{ Spam string }{}
				c.ReadJSON(&v)
				spam = v.Spam

				//New spam from user added to local spam list
				spamList = append(spamList, spam)
				appendToSpamDB(spam)
			}
			// os.Rename(fi, strings.Replace(fi, spam, "", 1))
		}
	}
	c.WriteJSON(&ResponseMsg{iter, cleaned_fi})
	c.Close()
}

func appendToSpamDB(sp string) {
	if sp != "" {
		file, _ := os.OpenFile("spam.txt", os.O_RDWR|os.O_APPEND, 0666)
		defer file.Close()
		b := make([]byte, 1000) //this can be efficient
		file.Read(b)
		if !strings.Contains(string(b), sp) {
			file.WriteString(sp + "\n")
		}
	}
}

func stringInSlice(a string, b []string) string {
	for _, el := range b {
		if strings.Contains(a, el) {
			return el
		}
	}
	return ""
}
