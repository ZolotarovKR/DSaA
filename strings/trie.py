class Node:
    def __init__(self) -> None:
        self.cs: list[Node | None] = [None] * 26
        self.size: int = 0

    def is_terminal(self) -> bool:
        return self.size - sum(c.size for c in self.cs if c is not None) > 0


class Trie:
    def __init__(self):
        self.root: Node = Node()

    def add(self, s: str) -> None:
        current = self.root

        for i in map(lambda c: ord(c) - ord("a"), s):
            current.size += 1

            if current.cs[i] is None:
                current.cs[i] = Node()

            current = current.cs[i]

        current.size += 1

    def size(self) -> int:
        return self.root.size

    def count(self, s: str) -> int:
        current = self.root

        for i in map(lambda c: ord(c) - ord("a"), s):
            if current.cs[i] is None:
                return 0

            current = current.cs[i]

        return current.size - sum(c.size for c in current.cs if c is not None)

    def contains(self, s: str) -> bool:
        return self.count(s) > 0

    def get(self, idx: int) -> str:
        chars: list[str] = []
        current = self.root

        while idx != 0 or not current.is_terminal():
            if current.is_terminal():
                idx -= 1

            for i, child in enumerate(current.cs):
                if child is None:
                    continue

                if idx < child.size:
                    chars.append(chr(i + ord("a")))
                    current = child
                    break

                idx -= child.size

        return "".join(chars)
