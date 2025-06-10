def log2(n: int) -> int:
    return n.bit_length() - 1


class SparseTable:
    def __init__(self, xs: list[int]) -> None:
        self.__st = [xs]
        for j in range(1, log2(len(xs)) + 1):
            self.__st.append([])
            for i in range(0, len(xs) - (1 << j) + 1):
                self.__st[j].append(
                    min(
                        self.__st[j - 1][i],
                        self.__st[j - 1][i + (1 << (j - 1))],
                    )
                )

    def query(self, l: int, r: int) -> int:
        assert r >= l
        t = log2(r - l + 1)
        return min(
            self.__st[t][l],
            self.__st[t][r - (1 << t) + 1],
        )
