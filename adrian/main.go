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
	fmt.Println("Load List")
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
	fmt.Println("Liste geladen")
	sort.Strings(lst)
	start := time.Now()
		
	//result := FindSubstring(lst, "Text") //simple contains suche
	result := BinarySearchSubstring(lst, "Text") //binäre Suche
	end := time.Now()
	fmt.Println(result)
	elapsed := end.Sub(start)
	fmt.Println("Zeit für die Suche:", elapsed, "Elemente gefunden:", len(result))

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

func FindSubstring(sliceStrings []string, substring string) []string {
    substring = strings.ToLower(substring)
    var matches []string
    for _, v := range sliceStrings {
	
        if strings.Contains(strings.ToLower(v), substring) {
            matches = append(matches, v)
        }
    }
    return matches
}



