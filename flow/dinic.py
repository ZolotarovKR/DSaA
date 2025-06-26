from linked_list import LinkedList
from collections import deque


def find_max_flow(n: int, cap: list[list[int]], s: int, t: int) -> int:
    lvl = [-1 for _ in range(n)]
    lvl[0] = 0

    def bfs() -> int:
        visited = [False for _ in range(n)]
        visited[s] = True
        queue = deque([s])

        while queue:
            u = queue.popleft()

            for v in range(n):
                if not visited[v] and cap[u][v] != 0:
                    lvl[v] = lvl[u] + 1
                    visited[v] = True
                    queue.append(v)

        return lvl[t]

    def dfs() -> int:
        visited = [False for _ in range(n)]
        visited[s] = True
        stack = [(LinkedList[int].singleton(s), (1 << 32) - 1)]

        while stack:
            path, flow = stack.pop()
            u = path.head().value

            if u == t:
                prev = path.head()
                while (curr := prev.ptr) is not None:
                    cap[prev.value][curr.value] += flow
                    cap[curr.value][prev.value] -= flow
                    prev = curr
                return flow

            for v in range(n):
                if not visited[v] and lvl[u] == lvl[v] - 1 and cap[u][v] != 0:
                    visited[v] = True
                    stack.append((path.cons(v), min(flow, cap[u][v])))
        return 0

    max_flow = 0
    while bfs() != -1:
        while (f := dfs()) != 0:
            max_flow += f
        lvl = [-1 for _ in range(n)]
        lvl[s] = 0
    return max_flow
