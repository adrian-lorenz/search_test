package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	var lst []string
	fmt.Println("Load list")
	start := time.Now()
	file, err := os.Open("wortliste.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		lst = append(lst, word)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		return
	}
	sort.Strings(lst)
	fmt.Println("list loaded")
	loadEnd := time.Now()
	start2 := time.Now()

	

	
	result := BinarySearchSubstring(lst, "Text") //binäre Suche

	end := time.Now()
	elapsed := end.Sub(start) //sum time
	elapsed2 := loadEnd.Sub(start) //load time
	elapsed3 := end.Sub(start2) //search time
	fmt.Println("Zeit für das Laden der Liste:", elapsed2)
	fmt.Println("Zeit für die Suche:", elapsed3)
	fmt.Println("Zeit für alles:", elapsed, "Elemente gefunden:", len(result))

}

func BinarySearchSubstring(sliceStrings []string, substring string) []string {
	matches := []string{}
	index := sort.Search(len(sliceStrings), func(i int) bool {
		return strings.Compare(sliceStrings[i], substring) >= 0
	})

	for i := index; i < len(sliceStrings); i++ {
		if strings.Contains(sliceStrings[i], substring) {
			matches = append(matches, sliceStrings[i])
		} else {
			break
		}
	}

	return matches
}

