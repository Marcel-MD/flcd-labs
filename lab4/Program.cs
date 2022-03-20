using lab4;

// Reading Context Free Grammar
var lines = File.ReadLines(@"C:\Users\Marcel\Projects\go\lfpc-labs\lab4\test2.txt").ToArray();

var n = lines[0].Split(" ").ToHashSet();
var t = lines[1].Split(" ").ToHashSet();
var p = new Dictionary<string, List<string>>();

foreach (var prod in lines[2].Split(" "))
{
    var tmp = prod.Split("-");

    if (!p.ContainsKey(tmp[0]))
    {
        p.Add(tmp[0], new List<string>() {tmp[1]});
    }
    else
    {
        p[tmp[0]].Add(tmp[1]);
    }
}
var grammar = new Grammar(){N = n, T = t, P = p, S = "S"};

Console.WriteLine("//== Step 0 ==//");
Console.WriteLine(grammar);

Console.WriteLine("//== Step 1 ==//");
Console.WriteLine(Step1.RemoveEmpty(grammar));

Console.WriteLine("//== Step 2 ==//");
Console.WriteLine(Step2.RemoveRename(grammar));

