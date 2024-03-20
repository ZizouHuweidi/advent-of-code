import sys


def Part1(s: list[str]) -> int:
    c = 0

    for ln in s:
        cr, lt, pw = ln.split(" ")
        lt = lt[0]
        min, max = [int(x) for x in cr.split("-")]

        if min <= pw.count(lt) <= max:
            c += 1

    return c


def Part2(s: list[str]) -> int:
    c = 0

    for ln in s:
        cr, lt, pw = ln.split(" ")
        lt = lt[0]
        min, max = [int(x) for x in cr.split("-")]

        if (pw[min - 1] == lt) ^ (pw[max - 1] == lt):  # xor
            c += 1

    return c


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")

    print(Part1(s))
    print(Part2(s))
