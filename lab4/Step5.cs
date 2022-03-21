namespace lab4;

public static class Step5
{
    public static Grammar Chomsky(Grammar grammar)
    {
        // {a, b} -> {X1, X2}
        var X = new Dictionary<string, string>();
        // {AB, SY1,...} -> {Y1, Y2,...}
        var Y = new Dictionary<string, string>();

        string GetX(string s)
        {
            if (X.TryGetValue(s, out var x))
                return x;

            x = "X" + (X.Count + 1);
            X.Add(s, x);
            return x;
        }
        
        string GetY(string s1, string s2)
        {
            var s = s1 + s2;
            if (Y.TryGetValue(s, out var y))
                return y;

            y = "Y" + (Y.Count + 1);
            Y.Add(s, y);
            return y;
        }

        string Convert(List<string> t)
        {
            if (t.Count == 1)
            {
                if (grammar.IsTerminal(t[0]))
                    return GetX(t[0]);
                return t[0];
            }
            
            return GetY(t[0], Convert(t.Skip(1).ToList()));
        }
        
        foreach (var (k, v) in grammar.P)
        {
            for (int i = 0; i < v.Count; i++)
            {
                if (v[i].Length < 2) continue;

                var t = new List<string>(v[i].Select(c => c.ToString()));

                if (grammar.IsTerminal(t[0]))
                {
                    t[0] = GetX(t[0]);
                }

                v[i] = t[0] + Convert(t.Skip(1).ToList());
            }
        }

        foreach (var (k, v) in X)
        {
            grammar.N.Add(v);
            grammar.P.Add(v, new List<string>{k});
        }
        
        foreach (var (k, v) in Y)
        {
            grammar.N.Add(v);
            grammar.P.Add(v, new List<string>{k});
        }

        return grammar;
    }
}