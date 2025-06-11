def z_function(s: str) -> list[int]:
    n = len(s)
    z = [0]
    l, r = 0, 0

    for i in range(1, n):
        z.append(min(r - i, z[i - l]) if i < r else 0)

        while i + z[i] < n and s[z[i]] == s[i + z[i]]:
            z[-1] += 1

        if i + z[-1] > r:
            l = i
            r = i + z[-1]

    return z


def find_matches(pat: str, txt: str):
    m = len(pat)
    n = len(txt)
    return (
        i
        for i, length in enumerate(z_function(pat + "#" + txt)[m + 1 : m + n - 1])
        if length == m
    )
