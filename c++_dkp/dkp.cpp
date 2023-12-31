#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <chrono>

#include "trie/trie.hpp"

int main(int, char **)
{
    std::cout << "Hello, from prefixSearch!\n";

    std::string line;
    std::vector<std::string> wortlist;

    std::ifstream myfile("wortliste.txt");
    if (myfile.is_open())
    {
        while (getline(myfile, line))
        {
            // std::cout << line << std::endl;
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

    trie testTrie;
    for (auto word : wortlist)
    {
        testTrie.add(word);
    }

    std::cout << "trie erstellt" << std::endl;
    auto start_time = std::chrono::high_resolution_clock::now();

    auto result = testTrie.Search("Text");
    auto end_time = std::chrono::high_resolution_clock::now();
    auto elapsed_time = std::chrono::duration_cast<std::chrono::microseconds>(end_time - start_time);

    std::cout << "Zeit für die Suche: " << elapsed_time.count() << " Mikrosekunden" << std::endl;

    /*
    for(auto entry : result) {
        std::cout << entry << std::endl;
    }*/
}