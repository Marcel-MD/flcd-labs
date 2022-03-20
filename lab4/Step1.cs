namespace lab4;

public static class Step1
{
    public static Grammar RemoveEmpty(Grammar grammar)
    {
        // Get the list of symbols that derive to ε
        var emptySymbols = new List<char>();
        foreach (var (k, v) in grammar.P)
        {
            if (v.Contains("ε"))
            {
                emptySymbols.Add(k[0]);
                grammar.P[k].Remove("ε");
            }
        }

        // Add new transitions
        foreach (var symbol in emptySymbols)
        {
            foreach (var (k, v) in grammar.P)
            {
                var length = v.Count;
                for (var i = 0; i < length; i++)
                {
                    var occurrence = v[i].Count(c => c == symbol);
                    if (occurrence > 0)
                    {
                        grammar.P[k].AddRange(GetNewTransitions(symbol, v[i], occurrence));
                    }
                }
            }   
        }

        return grammar;
    }

    private static List<string> GetNewTransitions(char c, string s, int occurrence)
    {
        if (occurrence > 2)
        {
            throw new NotImplementedException($"No combinatorics here, please change this '{s}'");
        }
        
        var ts = new List<string>();
    
        if (occurrence == 2)
        {
            ts.Add(s.Remove(s.IndexOf(c), 1));
            ts.Add(s.Remove(s.LastIndexOf(c), 1));
        }
        
        ts.Add(s.Replace(c.ToString(), ""));
        return ts;
    }
}