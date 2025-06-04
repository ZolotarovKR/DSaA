def log2(n: int) -> int:
    return n.bit_length() - 1


class BinaryLifting:
    def __init__(self, children: list[list[int]]) -> None:
        self.__n = len(children)
        self.__tin: list[int] = [0 for _ in range(self.__n)]
        self.__tout: list[int] = [0 for _ in range(self.__n)]
        self.__up: list[list[int]] = [[0] * log2(self.__n) for _ in range(self.__n)]

        def dfs(node: int, parent: int, time: int) -> int:
            time += 1
            self.__tin[node] = time
            self.__up[node][0] = parent
            for i in range(1, log2(self.__n)):
                self.__up[node][i] = self.__up[self.__up[node][i - 1]][i - 1]

            for child in children[node]:
                time = dfs(child, node, time)

            time += 1
            self.__tout[node] = time
            return time

        _ = dfs(0, 0, 0)

    def __is_ancestor(self, u: int, v: int) -> bool:
        return self.__tin[u] <= self.__tin[v] and self.__tout[u] >= self.__tout[v]

    def query(self, u: int, v: int) -> int:
        if self.__is_ancestor(u, v):
            return u
        if self.__is_ancestor(v, u):
            return v

        for i in range(log2(self.__n) - 1, -1, -1):
            if not self.__is_ancestor(self.__up[u][i], v):
                u = self.__up[u][i]
        return self.__up[u][0]
