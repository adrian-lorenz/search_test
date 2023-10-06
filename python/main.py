import time
from trie import Trie

def main():
    t = Trie()

    print("Load List")
    try:
        with open("wortliste.txt", "r") as file:
            for line in file:
                word = line.strip()
                t.add(word)
        print("Liste geladen")

        start = time.time()
        result = t.search("Text")
        end = time.time()

        print(result)
        
        elapsed_microseconds = (end - start) * 1000000
        print("Zeit für die Suche:", round(elapsed_microseconds,3), "Microseconds - Elemente gefunden:", len(result))

    except FileNotFoundError as e:
        print("Fehler beim Öffnen der Datei:", e)

if __name__ == "__main__":
    main()
