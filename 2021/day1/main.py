import sys


def Part1(input_list: list[int]) -> int:
    res = 0
    current = input_list[0]

    for input in input_list:
        if input > current:
            res += 1
        current = input

    return res


def Part2(input_list: list[int]) -> int:
    res = 0
    arr = []

    for i in range(len(input_list) - 2):
        arr.append(input_list[i] + input_list[i + 1] + input_list[i + 2])

    res = Part1(arr)

    return res


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read()

    input_list = list(map(int, s.strip().split("\n")))

    print(Part1(input_list))
    print(Part2(input_list))
