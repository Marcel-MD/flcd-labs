namespace lab4;

public class Step4
{
    public static Grammar RemoveInaccessible(Grammar grammar)
    {
        var accessible = new HashSet<string> {grammar.S};

        while (true)
        {
            var toAdd = new HashSet<string>();

            foreach (var symbol in accessible)
            {
                if (grammar.IsTerminal(symbol)) continue;
                
                foreach (var transition in grammar.P[symbol])
                {
                    foreach (var c in transition)
                    {
                        var s = c.ToString();
                        if (!accessible.Contains(s) && !toAdd.Contains(s))
                        {
                            toAdd.Add(s);
                        }
                    }
                }
            }
            
            if (toAdd.Count == 0) break;
            
            accessible.UnionWith(toAdd);
        }
        
        var inaccessible = grammar.N.Except(accessible).ToHashSet();
        
        // Remove inaccessible symbols
        grammar.RemoveEverythingContaining(inaccessible);
        
        return grammar;
    }
}