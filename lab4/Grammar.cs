using System.Text;

namespace lab4;

public class Grammar
{
    public List<string> N { get; set; }
    public List<string> T { get; set; }
    public Dictionary<string, List<string>> P { get; set; }
    public string S { get; set; }

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
        
        sb.Append("P = {\n");
        foreach (var (k, v) in P)
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