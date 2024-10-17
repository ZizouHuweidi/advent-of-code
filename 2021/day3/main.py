import sys


def Part1(input):
    N = len(input[0])

    epsilon = ["0"] * N
    gamma = ["0"] * N

    for i in range(N):
        zeros = sum(input[j][i] == "0" for j in range(len(input)))
        ones = sum(input[j][i] == "1" for j in range(len(input)))

        if zeros > ones:
            gamma[i] = "0"
            epsilon[i] = "1"
        else:
            gamma[i] = "1"
            epsilon[i] = "0"

    return int("".join(gamma), 2) * int("".join(epsilon), 2)


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        s = f.read().strip().split("\n")

    print(Part1(s))
