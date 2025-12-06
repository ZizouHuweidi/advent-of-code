using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

static long Part1(List<string> lines)
{
    const int Mod = 100;
    int pos = 50;
    long zeros = 0;

    foreach (var raw in lines)
    {
        var line = raw.Trim();
        if (line.Length < 2) continue;

        char dir = char.ToUpperInvariant(line[0]);

        if (!int.TryParse(line.Substring(1), out var dist))
            continue;

        if (dir == 'L')
        {
            pos = ((pos - dist) % Mod + Mod) % Mod;
        }
        else if (dir == 'R')
        {
            pos = (pos + dist) % Mod;
        }
        else
        {
            continue;
        }

        if (pos == 0) zeros++;
    }

    return zeros;
}

if (args.Length < 1)
{
    Console.WriteLine("Usage: dotnet run input.txt");
    return;
}

var filePath = args[0];

var lines = File.ReadLines(filePath)
                .Select(l => l.Trim())
                .Where(l => l.Length > 0)
                .ToList();

long part1 = Part1(lines);

Console.WriteLine($"Part 1: {part1}");
