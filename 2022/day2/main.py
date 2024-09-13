import sys

youScores = {"X": 1, "Y": 2, "Z": 3}
gameScores = {
    "AX": 3,
    "AY": 6,
    "AZ": 0,
    "BX": 0,
    "BY": 3,
    "BZ": 6,
    "CX": 6,
    "CY": 0,
    "CZ": 3,
}


def Part1(testInput):
    score = 0

    for line in testInput:
        plays = line.split(" ")
        oponent = plays[0]
        you = plays[1]
        score += youScores[you]
        score += gameScores[oponent + you]

    return score


winScores = {
    "X": 0,
    "Y": 3,
    "Z": 6,
}

chosenOutcomeScores = {
    "AX": 3,
    "AY": 1,
    "AZ": 2,
    "BX": 1,
    "BY": 2,
    "BZ": 3,
    "CX": 2,
    "CY": 3,
    "CZ": 1,
}


def Part2(testInput):
    score = 0
    for line in testInput:
        plays = line.split(" ")
        oponent = plays[0]
        win = plays[1]
        score += winScores[win]
        score += chosenOutcomeScores[oponent + win]

    return score


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")

    print(f"Part 1: {Part1(s)}")
    print(f"Part 2: {Part2(s)}")
