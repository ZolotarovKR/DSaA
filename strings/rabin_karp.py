P: int = 31
M: int = 1_610_612_741


def latin_small_to_int(c: str) -> int:
    return ord(c) - ord("a") + 1


def hash(s: str) -> int:
    h = 0
    for c in s:
        h = (h * P + latin_small_to_int(c)) % M
    h = h * P % M
    return h


def rabin_karp(pat: str, txt: str):
    m = len(pat)
    n = len(txt)
    h_pat = hash(pat)
    h_txt = hash(txt[0:m])

    if h_pat == h_txt:
        yield 0
    for i in range(1, n - m + 1):
        h_txt = (
            (
                h_txt
                - latin_small_to_int(txt[i - 1]) * P**m
                + latin_small_to_int(txt[i + m - 1])
            )
            * P
            % M
        )
        if h_pat == h_txt:
            yield i
