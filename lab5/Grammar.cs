using System.Text;

namespace lab5;

public class Grammar
{
    public HashSet<string> N { get; set; }
    public HashSet<string> T { get; set; }
    public Dictionary<string, HashSet<string>> P { get; set; }
    public string S { get; set; }
    
    private int _nextNonTerminal = 83;
    public Dictionary<string, HashSet<string>> First = new Dictionary<string, HashSet<string>>();
    public Dictionary<string, HashSet<string>> Follow = new Dictionary<string, HashSet<string>>();
    
    public bool IsTerminal(string symbol) => T.Contains(symbol);
    public bool IsTerminal(char symbol) => T.Contains(symbol.ToString());
    public bool IsNonTerminal(string symbol) => N.Contains(symbol);
    public bool IsNonTerminal(char symbol) => N.Contains(symbol.ToString());

    public string GetNextNonTerminal()
    {
        _nextNonTerminal++;
        var nt = ((char) _nextNonTerminal).ToString();
        N.Add(nt);
        return nt;
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
        
        sb.Append("First = {\n");
        foreach (var (k, v) in First)
        {
            foreach (var s in v)
            {
                sb.Append($"  {k} -> {s}\n");
            }
        }
        sb.Append("}\n");

        sb.Append("Follow = {\n");
        foreach (var (k, v) in Follow)
        {
            foreach (var s in v)
            {
                sb.Append($"  {k} -> {s}\n");
            }
        }
        sb.Append("}\n");
        
        return sb.ToString();
    }
}