namespace lab4;

public static class Step3
{
    public static Grammar RemoveUnproductive(Grammar grammar)
    {
        var productive = new HashSet<string>();

        // A -> α, where α in Vt
        foreach (var (k, v) in grammar.P)
        {
            foreach (var transition in v)
            {
                if (grammar.IsTerminal(transition))
                {
                    productive.Add(k);
                }
            }
        }
        
        // B -> b, where b in Vt or in productive
        foreach (var (k, v) in grammar.P)
        {
            if (!productive.Contains(k))
            {
                foreach (var transition in v)
                {
                    var count = transition.TakeWhile(symbol =>
                        grammar.IsTerminal(symbol) || productive.Contains(symbol.ToString())).Count();

                    if (transition.Length == count)
                    {
                        productive.Add(k);
                        break;
                    }
                }
            }
        }

        var unproductive = grammar.N.Except(productive).ToHashSet();

        // Remove unproductive
        foreach (var symbol in unproductive)
        {
            grammar.P.Remove(symbol);
        }
        
        foreach (var (k, v) in grammar.P)
        {
            var length = v.Count;
            for (var i = 0; i < length; i++)
            {
                bool containsUnproductive = v[i].Any(symbol => unproductive.Contains(symbol.ToString()));

                if (containsUnproductive)
                {
                    v.RemoveAt(i);
                    length--;
                    i--;
                }
            }
        }

        return grammar;
    }
}