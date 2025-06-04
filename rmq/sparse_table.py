def log2(n: int) -> int:
    return n.bit_length() - 1


class SparseTable:
    def __init__(self, xs: list[int]) -> None:
        self.__st = [xs]
        for k in range(1, log2(len(xs)) + 1):
            self.__st.append([])
            for i in range(0, len(xs) - (1 << k) + 1):
                self.__st[k].append(
                    min(
                        self.__st[k - 1][i],
                        self.__st[k - 1][i + (1 << (k - 1))],
                    )
                )

    def query(self, l: int, r: int) -> int:
        assert r >= l
        t = log2(r - l + 1)
        return min(
            self.__st[t][l],
            self.__st[t][r - (1 << t) + 1],
        )


def main() -> None:
    import random

    n = random.randint(17, 32)

    xs = list(random.choices(range(0, 256), k=n))
    st = SparseTable(xs)

    for l in range(len(xs)):
        for r in range(l, len(xs)):
            assert min(xs[l : r + 1]) == st.query(l, r)


if __name__ == "__main__":
    main()
