def build_z_function(s: str) -> list[int]:
    n = len(s)
    z = [n]
    l, r = 0, 0

    for i in range(1, n):
        z.append(min(r - i, z[i - l]) if i < r else 0)
        while i + z[i] < n and s[z[i]] == s[i + z[i]]:
            z[-1] += 1
        if i + z[-1] > r:
            l = i
            r = i + z[-1]

    return z


def z_function(pat: str, txt: str):
    return (
        i
        for i, length in enumerate(build_z_function(pat + "#" + txt))
        if length >= len(pat)
    )
