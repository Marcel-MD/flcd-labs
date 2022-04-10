namespace lab5;

public static class Step2
{
    public static void RemoveLeftFactoring(this Grammar grammar)
    {
        var haveLeftFactoring = new Dictionary<string, string>();

        while (true)
        {
            foreach (var (k, v) in grammar.P)
            {
                if(v.Count < 2) continue;
            
                string prefix = new string(
                    v.First().Substring(0, v.Min(s => s.Length))
                        .TakeWhile((c, i) => v.All(s => s[i] == c)).ToArray());

                if (prefix != "")
                {
                    haveLeftFactoring.Add(k, prefix);
                }
            }

            if (haveLeftFactoring.Count == 0) return;

            foreach (var (k, v) in haveLeftFactoring)
            {
                RemoveLeftFactoringFor(k, v);
            }
        
            haveLeftFactoring.Clear();
        }
        
        void RemoveLeftFactoringFor(string k, string prefix)
        {
            var nt = grammar.GetNextNonTerminal();
            var symbols = new List<string>();

            foreach (var s in grammar.P[k])
                if (s.StartsWith(prefix))
                    symbols.Add(s);
            
            grammar.P[k].Add(prefix + nt);
            grammar.P.Add(nt, new HashSet<string>());

            foreach (var s in symbols)
            {
                grammar.P[k].Remove(s);

                var suffix = s.Substring(prefix.Length);
                grammar.P[nt].Add(suffix != "" ? suffix : "ε");
            }
        }
    }
}