using System;
using System.Diagnostics;
using System.IO;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

public class Trie
{
    private class TrieNode
    {
        public Dictionary<char, TrieNode> Children { get; set; }
        public bool IsEndOfWord { get; set; }

        public TrieNode()
        {
            Children = new Dictionary<char, TrieNode>();
            IsEndOfWord = false;
        }
    }

    private readonly TrieNode root;

    public Trie()
    {
        root = new TrieNode();
    }

    public void Add(string word)
    {
        var currentNode = root;
        foreach (var letter in word)
        {
            if (!currentNode.Children.ContainsKey(letter))
            {
                currentNode.Children[letter] = new TrieNode();
            }
            currentNode = currentNode.Children[letter];
        }
        currentNode.IsEndOfWord = true;
    }

    public List<string> Search(string prefix)
    {
        var currentNode = root;
        foreach (var letter in prefix)
        {
            if (!currentNode.Children.ContainsKey(letter))
            {
                return new List<string>(); // Präfix existiert nicht
            }
            currentNode = currentNode.Children[letter];
        }
        return GetWordsFromNode(currentNode, new StringBuilder(prefix));
    }

    private List<string> GetWordsFromNode(TrieNode node, StringBuilder prefix)
    {
        var results = new List<string>();
        if (node.IsEndOfWord)
        {
            results.Add(prefix.ToString());
        }
        foreach (var child in node.Children)
        {
            prefix.Append(child.Key);
            results.AddRange(GetWordsFromNode(child.Value, prefix));
            prefix.Length--; // String zurücksetzen
        }
        return results;
    }
}

class Program
{
    static void Main()
    {
        var trie = new Trie();
        var stopwatch = new Stopwatch();
        Console.WriteLine("Load List");

        try
        {
            stopwatch.Start();

            // Lade die Datei parallel
            var lines = File.ReadAllLines("wortliste.txt");
            Parallel.ForEach(lines, line =>
            {
                var word = line.Trim();
                trie.Add(word);
            });

            stopwatch.Stop();
            Console.WriteLine("Liste geladen");
            var loadTime = stopwatch.Elapsed.TotalMilliseconds;

            // Suche starten
            stopwatch.Restart();
            var result = trie.Search("Text");
            stopwatch.Stop();
            var searchTime = stopwatch.Elapsed.TotalMilliseconds;

            // Gesamtlaufzeit
            var totalTime = loadTime + searchTime;

            Console.WriteLine($"Zeit für das Laden: {loadTime} Millisekunden");
            Console.WriteLine($"Zeit für die Suche: {searchTime} Millisekunden");
            Console.WriteLine($"Zeit für Alles: {totalTime} Millisekunden - Elemente gefunden: {result.Count}");
        }
        catch (FileNotFoundException e)
        {
            Console.WriteLine($"Fehler beim Öffnen der Datei: {e.Message}");
        }
    }
}
