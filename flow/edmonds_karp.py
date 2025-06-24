from collections import deque
from linked_list import LinkedList


def find_max_flow(n: int, cap: list[list[int]], s: int, t: int) -> int:
    visited = [False for _ in range(n)]

    def bfs() -> int:
        queue = deque([(LinkedList[int].singleton(s), (1 << 32) - 1)])

        while queue:
            path, flow = queue.popleft()
            u = path.head().value

            if u == t:
                prev = path.head()
                while (curr := prev.ptr) is not None:
                    cap[prev.value][curr.value] += flow
                    cap[curr.value][prev.value] -= flow
                    prev = curr
                return flow

            visited[u] = True

            for v in range(n):
                if not visited[v] and cap[u][v] != 0:
                    queue.append((path.cons(v), min(flow, cap[u][v])))
        return 0

    max_flow = 0
    while (f := bfs()) != 0:
        max_flow += f
        visited = [False for _ in range(n)]
    return max_flow
