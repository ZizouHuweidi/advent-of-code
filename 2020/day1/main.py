import sys


def Part1(input_list: list[int], target: int) -> int:
    input_set = set()
    for i in range(len(input_list)):
        if target - input_list[i] in input_set:
            return input_list[i] * (target - input_list[i])
        input_set.add(input_list[i])
    return -1


def Part2(input_list: list[int], target: int) -> int:
    input_set = set()
    for i in range(len(input_list)):
        for j in range(i + 1, len(input_list)):
            if target - input_list[i] - input_list[j] in input_set:
                return (
                    input_list[i]
                    * input_list[j]
                    * (target - input_list[i] - input_list[j])
                )

            input_set.add(input_list[j])

    return -1


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read()

    input_list = list(map(int, s.strip().split("\n")))
    target = 2020

    print(Part1(input_list, target))
    print(Part2(input_list, target))
