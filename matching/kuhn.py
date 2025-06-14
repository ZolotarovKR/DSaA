def find_maximum_matching(m: int, n: int, adj_list: list[list[int]]) -> list[tuple[int, int]]:
    mt: list[int] = [-1] * n

    def try_kuhn(u: int, used: list[bool]) -> bool:
        if used[u]:
            return False
        used[u] = True

        for v in adj_list[u]:
            if mt[v] == -1 or try_kuhn(mt[v], used):
                mt[v] = u
                return True

        return False

    for i in range(m):
        _ = try_kuhn(i, [False] * m)

    return [(u, v) for v, u in enumerate(mt) if u != -1]
