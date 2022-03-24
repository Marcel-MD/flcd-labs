namespace lab4;

public static class Step1
{
    public static Grammar RemoveEmpty(Grammar grammar)
    {
        while (true)
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

                if (v.Contains(""))
                {
                    emptySymbols.Add(k[0]);
                    grammar.P[k].Remove("");
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
                        if (v[i].Contains(symbol))
                        {
                            grammar.P[k].AddRange(Combinations(symbol, v[i]).Skip(1));
                        }
                    }
                }   
            }

            if (emptySymbols.Count == 0)
            {
                break;
            }

            emptySymbols.Clear();
        }

        return grammar;
    }

    private static IEnumerable<string> Combinations(char symbol, string str)
    {
        int firstSymbol = str.IndexOf(symbol);
        
        if (firstSymbol == -1) // Base case: no further combinations
            return new []{str};

        string prefix = str.Substring(0, firstSymbol);
        string suffix = str.Substring(firstSymbol + 1);
        
        // Recursion: Generate all combinations of suffix
        var recursiveCombinations = Combinations(symbol, suffix);

        // Return sequence in which each string is a concatenation of the
        // prefix, either symbol or empty, and one of the recursively-found suffixes
        return
            from s in new []{symbol.ToString(), ""}
            from recSuffix in recursiveCombinations
            select prefix + s + recSuffix;                                    
    }
}