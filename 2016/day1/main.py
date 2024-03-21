import sys


def Part1(s: list[str]) -> int:
    curr_dir = "n"
    x, y = 0, 0

    for d in s:
        turn = d[0]
        steps = int(d[1:])

        if turn == "R":
            if curr_dir == "n":
                x += steps
                curr_dir = "e"
            elif curr_dir == "s":
                x -= steps
                curr_dir = "w"
            elif curr_dir == "w":
                y += steps
                curr_dir = "n"
            elif curr_dir == "e":
                y -= steps
                curr_dir = "s"
        elif turn == "L":
            if curr_dir == "n":
                x -= steps
                curr_dir = "w"
            elif curr_dir == "s":
                x += steps
                curr_dir = "e"
            elif curr_dir == "w":
                y -= steps
                curr_dir = "s"
            elif curr_dir == "e":
                y += steps
                curr_dir = "n"

    return abs(x) + abs(y)


def Part2(s: list[str]) -> int:
    curr_dir = "n"
    x, y = 0, 0
    visited = set()
    visited.add((0, 0))

    for d in s:
        turn = d[0]
        steps = int(d[1:])

        if turn == "R":
            if curr_dir == "n":
                for _ in range(steps):
                    x += 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "e"
            elif curr_dir == "s":
                for _ in range(steps):
                    x -= 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "w"
            elif curr_dir == "w":
                for _ in range(steps):
                    y += 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "n"
            elif curr_dir == "e":
                for _ in range(steps):
                    y -= 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "s"
        elif turn == "L":
            if curr_dir == "n":
                for _ in range(steps):
                    x -= 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "w"
            elif curr_dir == "s":
                for _ in range(steps):
                    x += 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "e"
            elif curr_dir == "w":
                for _ in range(steps):
                    y -= 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "s"
            elif curr_dir == "e":
                for _ in range(steps):
                    y += 1
                    if (x, y) in visited:
                        return abs(x) + abs(y)
                    visited.add((x, y))
                curr_dir = "n"

    return -1


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split(", ")

    print(Part1(s))
    print(Part2(s))
