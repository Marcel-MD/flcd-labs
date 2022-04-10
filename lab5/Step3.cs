namespace lab5;

public static class Step3
{
    public static void FirstFollow(this Grammar g)
    {
        //== First ==//
        
        foreach (var (k, v) in g.P)
        {
            g.First.Add(k, new HashSet<string>());
            First(k, g.First[k]);
        }
        
        void First(string symbol, HashSet<string> set)
        {
            var fs = symbol.First().ToString();
            if (g.IsTerminal(fs))
            {
                set.Add(fs);
                return;
            }
            
            foreach (var s in g.P[fs])
            {
                First(s, set);
            }
        }
        
        //== Follow ==//
        
        foreach (var (k, v) in g.P)
        {
            g.Follow.Add(k, new HashSet<string>());
        }
        g.Follow[g.S].Add("$");

        foreach (var (k, v) in g.P)
        {
            foreach (var s in v)
            {
                for (int i = 0; i < s.Length; i++)
                {
                    if (g.IsNonTerminal(s[i]))
                    {
                        g.Follow[s[i].ToString()].UnionWith(GetFollow(k, i + 1, s));
                    }
                }
            }
        }

        HashSet<string> GetFollow(string k, int i, string s)
        {
            if (i >= s.Length)
                return g.Follow[k];

            var nt = s[i].ToString();

            if (g.IsTerminal(nt))
                return new HashSet<string>{nt};

            if (g.IsNonTerminal(nt))
            {
                var set = new HashSet<string>();
                set.UnionWith(g.First[nt]);
                if (g.First[nt].Contains("ε"))
                    set.UnionWith(GetFollow(k, i + 1, s));
                set.Remove("ε");
                return set;
            }

            return new HashSet<string>();
        }
    }
}