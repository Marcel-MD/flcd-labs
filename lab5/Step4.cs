namespace lab5;

public static class Step4
{
    public static void ConstructParsingTable(this Grammar g)
    {
        var table = new Dictionary<string, Dictionary<string, string>>();

        foreach (var n in g.N)
        {
            table.Add(n, new Dictionary<string, string>());
            foreach (var t in g.T)
                table[n].Add(t, "");
            table[n].Remove("ε");
            table[n].Add("$", "");
        }

        foreach (var (k, v) in g.P)
        {
            foreach (var s in v)
            {
                if (s == "ε")
                {
                    foreach (var follow in g.Follow[k])
                    {
                        table[k][follow] = s;
                    }
                    continue;
                }
                
                var fs = s.First().ToString();
                if (g.First[k].Contains(fs))
                {
                    table[k][fs] = s;
                }
                else
                {
                    foreach (var first in g.First[k])
                    {
                        if (first == "ε") continue;
                        
                        if (table[k][first] == "")
                            table[k][first] = s;
                    }
                }
            }
        }

        g.ParsingTable = table;
    }
}