import sys
import re


def Part1(input_str):
    game_subset = input_str.split(":")[1].split(";")
    pattern = re.compile(r"(\d+)\s+([a-zA-Z]+)")
    for val in game_subset:
        matches = pattern.findall(val)
        for number, color in matches:
            number = int(number)
            if (
                (color == "red" and number > 12)
                or (color == "green" and number > 13)
                or (color == "blue" and number > 14)
            ):
                return False
    return True


def Part2(input_str):
    game_subset = input_str.split(":")[1].split(";")
    red, blue, green = 0, 0, 0
    pattern = re.compile(r"(\d+)\s+([a-zA-Z]+)")
    for val in game_subset:
        matches = pattern.findall(val)
        for number, color in matches:
            number = int(number)
            if color == "red":
                red = max(red, number)
            elif color == "blue":
                blue = max(blue, number)
            elif color == "green":
                green = max(green, number)
    return red * blue * green


if __name__ == "__main__":
    with open(sys.argv[1], "r") as file:
        res_part1 = 0
        res_part2 = 0
        for line_number, line in enumerate(file, 1):
            if Part1(line.strip()):
                res_part1 += line_number
            power_number = Part2(line.strip())
            res_part2 += power_number

    print(f"Part_1: {res_part1}")
    print(f"Part_2: {res_part2}")
