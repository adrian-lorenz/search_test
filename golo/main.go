package main

import (
	"bufio"
	"fmt"
	"golo/trie"
	"os"
	"strings"
	"time"
)

func main() {
	t := trie.NewTrie()

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
		t.Add(word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		return
	}
	fmt.Println("Liste geladen")
	start := time.Now()
	result := t.Search("Text")
	end := time.Now()
	fmt.Println(result)
	elapsed := end.Sub(start)
	fmt.Println("Zeit für die Suche:", elapsed, "Elemente gefunden:", len(result))

}
