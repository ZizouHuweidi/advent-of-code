import sys

number_map = {
    "zero": "zero0zero",
    "one": "one1one",
    "two": "two2two",
    "three": "three3three",
    "four": "four4four",
    "five": "five5five",
    "six": "six6six",
    "seven": "seven7seven",
    "eight": "eight8eight",
    "nine": "nine9nine",
}


def Part1(line):
    digits = [int(char) for char in line if char.isdigit()]
    if digits:
        return digits[0] * 10 + digits[-1]
    return 0


def Part2(line):
    for word, digit in number_map.items():
        line = line.replace(word, digit)
    return Part1(line)


if __name__ == "__main__":
    with open(sys.argv[1], "r") as file:
        sum_part1 = 0
        sum_part2 = 0
        for line in file:
            combined_digits = Part1(line.strip())
            sum_part1 += combined_digits
            combined_digits = Part2(line.strip())
            sum_part2 += combined_digits

    print(f"Sum of Part_1: {sum_part1}")
    print(f"Sum of part_2: {sum_part2}")
