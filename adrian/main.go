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
	start := time.Now()
	sort.Strings(lst)
	
	result := findSubstring(lst, "apfel")
	end := time.Now()
	fmt.Println(result)
	elapsed := end.Sub(start)
	fmt.Printf("Zeit für die Suche: %s\n", elapsed)

}

func findSubstring(sliceStrings []string, substring string) []string {
    substring = strings.ToLower(substring)
    var matches []string
    for _, v := range sliceStrings {
        if strings.Contains(strings.ToLower(v), substring) {
            matches = append(matches, v)
        }
    }
    return matches
}



