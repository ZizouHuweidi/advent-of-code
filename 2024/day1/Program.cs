using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

static long Part1(List<long> left, List<long> right)
{
    var leftSorted = left.OrderBy(v => v).ToList();
    var rightSorted = right.OrderBy(v => v).ToList();

    return leftSorted.Zip(rightSorted, (a, b) => Math.Abs(a - b)).Sum();
}

static long Part2(List<long> left, List<long> right)
{
    var rightCounts = right
        .GroupBy(v => v)
        .ToDictionary(g => g.Key, g => g.Count());

    long similarity = 0;
    foreach (var value in left)
    {
        if (rightCounts.TryGetValue(value, out int count))
        {
            similarity += value * (long)count;
        }
    }

    return similarity;
}

if (args.Length < 1)
{
    Console.WriteLine("Usage: dotnet run input.txt");
    return;
}

var filePath = args[0];

var lines = File.ReadLines(filePath)
                .Select(l => l.Trim())
                .Where(l => !string.IsNullOrWhiteSpace(l))
                .ToList();

var left = new List<long>();
var right = new List<long>();

foreach (var line in lines)
{
    var parts = line.Split((char[])null, StringSplitOptions.RemoveEmptyEntries);
    if (parts.Length < 2) continue;

    left.Add(long.Parse(parts[0]));
    right.Add(long.Parse(parts[1]));
}

long sumPart1 = Part1(left, right);
long sumPart2 = Part2(left, right);

Console.WriteLine($"Part 1: {sumPart1}");
Console.WriteLine($"Part 2: {sumPart2}");
