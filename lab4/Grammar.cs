using System.Text;

namespace lab4;

public class Grammar
{
    public HashSet<string> N { get; set; }
    public HashSet<string> T { get; set; }
    public Dictionary<string, List<string>> P { get; set; }
    public string S { get; set; }

    public bool IsTerminal(string symbol) => T.Contains(symbol);
    public bool IsTerminal(char symbol) => T.Contains(symbol.ToString());
    public bool IsNonTerminal(string symbol) => N.Contains(symbol);
    public bool IsNonTerminal(char symbol) => N.Contains(symbol.ToString());

    public void RemoveEverythingContaining(HashSet<string> symbolSet)
    {
        foreach (var symbol in symbolSet)
        {
            P.Remove(symbol);
        }
        
        foreach (var (k, v) in P)
        {
            var length = v.Count;
            for (var i = 0; i < length; i++)
            {
                var containsSymbol = v[i].Any(symbol => symbolSet.Contains(symbol.ToString()));

                if (containsSymbol)
                {
                    v.RemoveAt(i);
                    length--;
                    i--;
                }
            }
        }
        
        N.ExceptWith(symbolSet);
        T.ExceptWith(symbolSet);
    }

    public override string ToString()
    {
        var sb = new StringBuilder();

        sb.Append("\nVn = { ");
        foreach (var s in N)
        {
            sb.Append($"{s} ");
        }
        sb.Append("}\n");
        
        sb.Append("Vt = { ");
        foreach (var s in T)
        {
            sb.Append($"{s} ");
        }
        sb.Append("}\n");

        var length = 0;
        sb.Append("P = {\n");
        foreach (var (k, v) in P)
        {
            foreach (var s in v)
            {
                sb.Append($"  {k} -> {s}\n");
                length++;
            }
        }
        sb.Append("}\nP = " + length + "\n");

        return sb.ToString();
    }
}