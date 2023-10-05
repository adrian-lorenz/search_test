package trie

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Add(word string) {
	currentNode := t.root
	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = NewNode()
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isEndOfWord = true
}
func (t *Trie) Search(prefix string) []string {
	currentNode := t.root
	for _, char := range prefix {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = NewNode()
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.GetWords(prefix)
}

// Vergeblicher Versuch einer rekursiven Lösung (contains) was aber mit der Datenstuktur nicht funktioniert
func (t *Trie) Search2(prefix string) []string {
	currentNode := t.root
	result := []string{}

	var dfs func(node *Node, word string)
	dfs = func(node *Node, word string) {
		if node.isEndOfWord {
			result = append(result, word)
		}
		for char, child := range node.children {
			dfs(child, word+string(char))
		}
	}

	for _, char := range prefix {
		if _, ok := currentNode.children[char]; !ok {
			return result
		}
		currentNode = currentNode.children[char]
	}

	dfs(currentNode, prefix) // Starten der DFS-Suche ab dem aktuellen Knoten

	return result
}
