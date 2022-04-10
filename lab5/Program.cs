// Reading Context Free Grammar

using lab5;

var lines = File.ReadLines(@"C:\Users\Marcel\Projects\go\lfpc-labs\lab5\test2.txt").ToArray();

var n = lines[0].Split(" ").ToHashSet();
var t = lines[1].Split(" ").ToHashSet();
var p = new Dictionary<string, HashSet<string>>();

foreach (var prod in lines[2].Split(" "))
{
    var tmp = prod.Split("-");

    if (!p.ContainsKey(tmp[0]))
    {
        p.Add(tmp[0], new HashSet<string>() {tmp[1]});
    }
    else
    {
        p[tmp[0]].Add(tmp[1]);
    }
}
var grammar = new Grammar(){N = n, T = t, P = p, S = "S"};

Console.WriteLine("//== Default Grammar ==//");
Console.WriteLine(grammar);

Console.WriteLine("//== Remove Left Recursion ==//");
grammar.RemoveLeftRecursion();
Console.WriteLine(grammar);

Console.WriteLine("//== Remove Left Factoring ==//");
grammar.RemoveLeftFactoring();
Console.WriteLine(grammar);

Console.WriteLine("//== First and Follow ==//");
grammar.FirstFollow();
grammar.PrintFirstFollow();

Console.WriteLine("//== Parsing Table ==//");
grammar.ConstructParsingTable();
grammar.PrintParsingTable();