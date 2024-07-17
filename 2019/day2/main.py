import sys


def Part1(s: list[int]) -> int:
    for i in range(0, len(s), 4):
        opcode = s[i]
        if opcode == 99:
            break
        if opcode == 1:
            s[s[i + 3]] = s[s[i + 1]] + s[s[i + 2]]
        elif opcode == 2:
            s[s[i + 3]] = s[s[i + 1]] * s[s[i + 2]]
    return s[0]


# def Part2()


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split(",")

    s = list(map(int, s))
    s[1] = 12
    s[2] = 2
    print(Part1(s))
