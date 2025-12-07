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

static long Part2(List<string> lines)
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

        long d = dist;

        int r = pos % Mod;
        if (r < 0) r += Mod;

        int target_k = dir == 'R' ? ((Mod - r) % Mod) : (r % Mod);
        if (target_k == 0) target_k = Mod;

        long hits = 0;
        if (d >= target_k)
        {
            hits = 1 + (d - target_k) / Mod;
        }

        zeros += hits;

        int step = dist % Mod;
        if (dir == 'L')
        {
            pos = ((pos - step) % Mod + Mod) % Mod;
        }
        else if (dir == 'R')
        {
            pos = (pos + step) % Mod;
        }
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
long part2 = Part2(lines);

Console.WriteLine($"Part 1: {part1}");
Console.WriteLine($"Part 2: {part2}");
