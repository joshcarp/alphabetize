package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
)
var re = regexp.MustCompile(`[^\s][^{]*{[^}]*}`)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	matches := re.FindAllString(string(bytes), -1)
	sort.Strings(matches)
	output, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	for i, match := range matches {
		if i != len(matches) -1{
			match = match + "\n"
		}
		_, err := output.WriteString(match + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}