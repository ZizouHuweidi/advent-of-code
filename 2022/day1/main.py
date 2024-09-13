import sys


def Part1(testInput):
    elfTotals = []
    currentTotal = 0
    for line in testInput:
        if line == "":
            elfTotals.append(currentTotal)
            currentTotal = 0
        else:
            cal = int(line)
            currentTotal += cal

    if currentTotal > 0:
        elfTotals.append(currentTotal)

    elfTotals.sort(reverse=True)
    return elfTotals[:3]


def Part2(nums):
    return sum(nums)


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")
    top_three = Part1(s)
    print(f"Part 1: {top_three}")
    print(f"Part 2: {Part2(top_three)}")
