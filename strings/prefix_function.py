def prefix_function(s: str) -> list[int]:
    n = len(s)
    pi = [0]

    for i in range(1, n):
        j = pi[i - 1]

        while j > 0 and s[i] != s[j]:
            j = pi[j - 1]

        pi.append(j + 1 if s[i] == s[j] else j)

    return pi


def find_matches(pat: str, txt: str):
    m = len(pat)
    return (
        i - m + 1
        for i, length in enumerate(prefix_function(pat + "#" + txt)[m + 1 :])
        if length == m
    )
