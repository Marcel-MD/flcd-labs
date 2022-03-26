namespace lab4;

public static class Step3
{
    public static Grammar RemoveUnproductive(Grammar grammar)
    {
        var productive = new HashSet<string>();

        // A -> αa, where α in Vt
        foreach (var (k, v) in grammar.P)
        {
            foreach (var transition in v)
            {
                var count = transition.TakeWhile(c => grammar.IsTerminal(c.ToString())).Count();

                if (transition.Length == count)
                {
                    productive.Add(k);
                }
            }
        }
        
        // B -> b, where b in Vt or in productive
        while (true)
        {
            var toAdd = new HashSet<string>();
            
            foreach (var (k, v) in grammar.P)
            {
                if (!productive.Contains(k) && !toAdd.Contains(k))
                {
                    foreach (var transition in v)
                    {
                        var count = transition.TakeWhile(symbol =>
                            grammar.IsTerminal(symbol) ||
                            productive.Contains(symbol.ToString()) ||
                            toAdd.Contains(symbol.ToString())).Count();

                        if (transition.Length == count)
                        {
                            toAdd.Add(k);
                            break;
                        }
                    }
                }
            }
            
            if (toAdd.Count == 0) break;
            
            productive.UnionWith(toAdd);
        }

        var unproductive = grammar.N.Except(productive).ToHashSet();

        // Remove unproductive symbols
        grammar.RemoveEverythingContaining(unproductive);

        return grammar;
    }
}