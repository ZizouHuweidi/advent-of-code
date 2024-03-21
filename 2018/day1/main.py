import sys


def Part1(s: list[str]) -> int:
    res = 0

    for x in s:
        sign, amount = x[0], int(x[1:])
        if sign == "+":
            res += amount
        if sign == "-":
            res -= amount

    return res


def Part2(s: list[str]) -> int:
    frequencies = set()
    current_frequency = 0

    while True:
        for x in s:
            sign, amount = x[0], int(x[1:])
            if sign == "+":
                current_frequency += amount
            elif sign == "-":
                current_frequency -= amount

            if current_frequency in frequencies:
                return current_frequency
            else:
                frequencies.add(current_frequency)


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")

    print(Part1(s))
    print(Part2(s))
