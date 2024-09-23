const fs = require('fs');
const { performance } = require('perf_hooks');

class TrieNode {
  constructor() {
    this.children = {};
    this.isEndOfWord = false;
  }
}

class Trie {
  constructor() {
    this.root = new TrieNode();
  }

  add(word) {
    let currentNode = this.root;
    for (let char of word) {
      if (!currentNode.children[char]) {
        currentNode.children[char] = new TrieNode();
      }
      currentNode = currentNode.children[char];
    }
    currentNode.isEndOfWord = true;
  }

  search(prefix) {
    let currentNode = this.root;
    for (let char of prefix) {
      if (!currentNode.children[char]) {
        return []; // Prefix not found
      }
      currentNode = currentNode.children[char];
    }
    return this._collectWordsFromNode(currentNode, prefix);
  }

  _collectWordsFromNode(node, prefix) {
    let results = [];
    if (node.isEndOfWord) {
      results.push(prefix);
    }
    for (let char in node.children) {
      results.push(...this._collectWordsFromNode(node.children[char], prefix + char));
    }
    return results;
  }
}

function main() {
  const trie = new Trie();

  // Start loading time
  const startLoadTime = performance.now();

  console.log("Load List");

  try {
    const data = fs.readFileSync('wortliste.txt', 'utf-8');
    const lines = data.split('\n');

    lines.forEach(line => {
      const word = line.trim();
      trie.add(word);
    });

    console.log("Liste geladen");

    const endLoadTime = performance.now();
    const loadTime = endLoadTime - startLoadTime;

    // Start search time
    const startSearchTime = performance.now();
    const result = trie.search('Text');
    const endSearchTime = performance.now();
    const searchTime = endSearchTime - startSearchTime;

    // Total time
    const totalTime = endSearchTime - startLoadTime;

    console.log(`Zeit für das Laden: ${loadTime.toFixed(3)} Millisekunden`);
    console.log(`Zeit für die Suche: ${searchTime.toFixed(3)} Millisekunden`);
    console.log(`Zeit für Alles: ${totalTime.toFixed(3)} Millisekunden - Elemente gefunden: ${result.length}`);

  } catch (err) {
    console.error("Fehler beim Öffnen der Datei:", err.message);
  }
}

main();
