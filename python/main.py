import time
from trie import Trie

def main():
    t = Trie()
    start = time.time()
    print("Load List")
    try:
        with open("wortliste.txt", "r") as file:
            for line in file:
                word = line.strip()
                t.add(word)
        print("Liste geladen")
        load_end = time.time()
        start2 = time.time()
       
        result = t.search("Text")
        end = time.time()
       
        
        elapsed_microseconds = (end - start) * 1000000 #alles
        elapsed_microseconds2 = (end - start2) * 1000000 #suche
        elapsed_microseconds3 = (load_end - start) * 1000000 #laden
        print("Zeit für das Laden:", round(elapsed_microseconds3,3), "Microseconds")
        print("Zeit für die Suche:", round(elapsed_microseconds2,3), "Microseconds")
        print("Zeit für die Alles:", round(elapsed_microseconds,3), "Microseconds - Elemente gefunden:", len(result))

    except FileNotFoundError as e:
        print("Fehler beim Öffnen der Datei:", e)

if __name__ == "__main__":
    main()
