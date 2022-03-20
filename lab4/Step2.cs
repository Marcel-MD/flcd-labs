namespace lab4;

public static class Step2
{
    public static Grammar RemoveRename(Grammar grammar)
    {
        foreach (var (k, v) in grammar.P)
        {
            var length = v.Count;
            for (int i = 0; i < length; i++)
            {
                if (v[i].Length == 1 && grammar.IsNonTerminal(v[i]))
                {
                    var symbol = v[i];
                    foreach (var transition in grammar.P[symbol])
                    {
                        if (transition.Length > 1 || grammar.IsTerminal(transition))
                        {
                            if (v[i] == symbol)
                            {
                                v[i] = transition;
                            }
                            else
                            {
                                v.Add(transition);
                            }
                        }
                    }
                }
            }
        }

        return grammar;
    }
}