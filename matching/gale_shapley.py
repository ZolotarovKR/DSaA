def find_stable_matching(
    n: int, l_prefs: list[list[int]], r_prefs: list[list[int]]
) -> list[tuple[int, int]]:
    mt = [-1] * n
    used = [False] * n

    for i in range(n):
        for u in range(n):
            v = l_prefs[u][i]

            if used[u] or (
                mt[v] != -1 and r_prefs[v].index(mt[v]) < r_prefs[v].index(u)
            ):
                continue

            used[mt[v]] = False
            used[u] = True
            mt[v] = u

    return [(u, v) for v, u in enumerate(mt)]
