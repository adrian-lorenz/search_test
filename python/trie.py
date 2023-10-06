class Node:
    def __init__(self):
        self.children = {}
        self.is_end_of_word = False

    def get_words(self, prefix=''):
        result = []
        if self.is_end_of_word:
            result.append(prefix)
        for char, child_node in self.children.items():
            result.extend(child_node.get_words(prefix + char))
        return result



class Trie:
    def __init__(self):
        self.root = Node()

    def add(self, word):
        current_node = self.root
        for char in word:
            if char not in current_node.children:
                current_node.children[char] = Node()
            current_node = current_node.children[char]
        current_node.is_end_of_word = True

    def search(self, prefix):
        current_node = self.root
        for char in prefix:
            if char not in current_node.children:
                return []
            current_node = current_node.children[char]
        return current_node.get_words(prefix)
