from linked_list import LinkedList


def find_max_flow(n: int, cap: list[list[int]], s: int, t: int) -> int:
    visited = [False for _ in range(n)]

    def dfs() -> int:
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

            visited[u] = True

            for v in range(n):
                if not visited[v] and cap[u][v] != 0:
                    stack.append((path.cons(v), min(flow, cap[u][v])))
        return 0

    max_flow = 0
    while (f := dfs()) != 0:
        max_flow += f
        visited = [False for _ in range(n)]
    return max_flow
