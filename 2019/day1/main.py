import sys


def Part1(numArr):
    sum = 0

    for v in numArr:
        sum += int(v / 3) - 2

    return sum


def Part2(numArr):
    sum = 0

    for v in numArr:
        sum += calculateFuel(v)

    return sum


def calculateFuel(mass):
    fuel = (mass // 3) - 2

    if fuel <= 0:
        return 0

    return fuel + calculateFuel(fuel)


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")

    numArr = []

    for v in s:
        numArr.append(int(v))

    print(Part1(numArr))
    print(Part2(numArr))
