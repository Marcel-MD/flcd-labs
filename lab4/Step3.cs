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

        // Remove unproductive symbols
        grammar.RemoveEverythingContaining(unproductive);

        return grammar;
    }
}