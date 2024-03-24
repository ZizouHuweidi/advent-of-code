import sys


def Part1(s):
    res = ""
    key = 5

    for line in s:
        for c in line:
            if c == "U" and key < 4:
                continue
            if c == "D" and key > 6:
                continue
            if c == "R" and key in [3, 6, 9]:
                continue
            if c == "L" and key in [1, 4, 7]:
                continue

            if c == "U":
                key -= 3
            elif c == "D":
                key += 3
            elif c == "R":
                key += 1
            elif c == "L":
                key -= 1

        res += str(key)

    return res


def Part2(s):
    res = ""
    row, col = 3, 1

    phonepad = [
        [" ", " ", " ", " ", " ", " ", " "],
        [" ", " ", " ", "1", " ", " ", " "],
        [" ", " ", "2", "3", "4", " ", " "],
        [" ", "5", "6", "7", "8", "9", " "],
        [" ", " ", "A", "B", "C", " ", " "],
        [" ", " ", " ", "D", " ", " ", " "],
        [" ", " ", " ", " ", " ", " ", " "],
    ]

    for line in s:
        for c in line:
            if c == "U":
                if phonepad[row - 1][col] != " ":
                    row -= 1
            elif c == "D":
                if phonepad[row + 1][col] != " ":
                    row += 1
            elif c == "R":
                if phonepad[row][col + 1] != " ":
                    col += 1
            elif c == "L":
                if phonepad[row][col - 1] != " ":
                    col -= 1

        res += phonepad[row][col]

    return res


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")

    print(Part1(s))
    print(Part2(s))
