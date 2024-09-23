#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <unordered_map>
#include <chrono>

class TrieNode
{
public:
    std::unordered_map<char, TrieNode *> children;
    bool isEndOfWord;

    TrieNode()
    {
        isEndOfWord = false;
    }
};

class PrefixTrie
{
public:
    TrieNode *root;

    PrefixTrie()
    {
        root = new TrieNode();
    }

    void insert(const std::string &word)
    {
        TrieNode *current = root;
        for (char c : word)
        {
            if (current->children.find(c) == current->children.end())
            {
                current->children[c] = new TrieNode();
            }
            current = current->children[c];
        }
        current->isEndOfWord = true;
    }

    std::vector<std::string> search(const std::string &prefix)
    {
        std::vector<std::string> result;
        TrieNode *current = root;
        for (char c : prefix)
        {
            if (current->children.find(c) == current->children.end())
            {
                return result; // Prefix not found
            }
            current = current->children[c];
        }
        findWordsFromNode(current, prefix, result);
        return result;
    }

private:
    void findWordsFromNode(TrieNode *node, std::string currentWord, std::vector<std::string> &result)
    {
        if (node->isEndOfWord)
        {
            result.push_back(currentWord);
        }
        for (const auto &pair : node->children)
        {
            findWordsFromNode(pair.second, currentWord + pair.first, result);
        }
    }
};

int main()
{
    auto loadstart_time = std::chrono::high_resolution_clock::now();
    PrefixTrie trie;
    std::cout << "Hello, from prefixSearch!\n";

    std::string line;
    std::vector<std::string> wortlist;

    std::ifstream myfile("wortliste.txt");
    if (myfile.is_open())
    {
        while (getline(myfile, line))
        {
            wortlist.push_back(line);
        }
        myfile.close();
    }
    else
    {
        std::cout << "Unable to open file" << std::endl;
        return 99;
    }

    std::cout << wortlist.size() << " words loaded" << std::endl;

    for (auto word : wortlist)
    {
        trie.insert(word);
    }

    std::cout << "trie erstellt" << std::endl;
    auto endload_time = std::chrono::high_resolution_clock::now();
    auto start_time = std::chrono::high_resolution_clock::now();
    std::vector<std::string> result = trie.search("Text");
    auto end_time = std::chrono::high_resolution_clock::now();
    auto elapsed_time = std::chrono::duration_cast<std::chrono::microseconds>(end_time - start_time);
    auto elapsed_loadtime = std::chrono::duration_cast<std::chrono::microseconds>(endload_time - loadstart_time);
    std::cout << "Zeit für das Laden: " << elapsed_loadtime.count() << " Mikrosekunden" << std::endl;
    std::cout << "Zeit für die Suche: " << elapsed_time.count() << " Mikrosekunden" << std::endl;
  

    return 0;
}
