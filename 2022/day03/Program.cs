using System.Linq;

var input = File.ReadAllText("input.txt");

var part1 = input.Lines
    .Select(x => x.Batch(x.Length / 2))
    .Select(x => x.First().Intersect(x.Last()))
    .SelectMany(x => x)
    .Select(x => char.IsLower(x) ? x - 'a' + 1 : x - 'A' + 27)
    .Sum()
    .ToString();





