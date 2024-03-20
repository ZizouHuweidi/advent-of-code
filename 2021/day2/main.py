import sys


def Part1(input_list: list[str]) -> int:
    x = 0
    y = 0

    for line in input_list:
        direction, distance = line.split(" ")
        distance = int(distance)
        if direction == "forward":
            x += distance
        elif direction == "down":
            y += distance
        elif direction == "up":
            y -= distance
    return x * y


def Part2(input_list: list[str]) -> int:
    x = 0
    y = 0
    aim = 0

    for line in input_list:
        direction, distance = line.split(" ")
        distance = int(distance)
        if direction == "forward":
            x += distance
            y += aim * distance
        elif direction == "down":
            aim += distance
        elif direction == "up":
            aim -= distance
    return x * y


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")

    print(Part1(s))
    print(Part2(s))
