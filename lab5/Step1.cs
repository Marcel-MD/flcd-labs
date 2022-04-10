namespace lab5;

public static class Step1
{
    public static void RemoveLeftRecursion(this Grammar grammar)
    {
        var haveLeftRecursion = new HashSet<string>();
        
        foreach (var (k, v) in grammar.P)
        {
            foreach (var s in v)
            {
                if (k == s.First().ToString())
                {
                    haveLeftRecursion.Add(k);
                    break;
                }
            }
        }

        foreach (var s in haveLeftRecursion)
        {
            removeLeftRecursionFor(k: s);
        }
        
        void removeLeftRecursionFor(string k)
        {
            var nt = grammar.GetNextNonTerminal();

            var alfa = new List<string>();
            var beta = new List<string>();

            foreach (var s in grammar.P[k])
            {
                if (k == s.First().ToString())
                {
                    alfa.Add(s.Substring(1));
                }
                else
                {
                    beta.Add(s);
                }
            }
            
            grammar.P[k].Clear();
            foreach (var s in beta)
            {
                grammar.P[k].Add(s);
                grammar.P[k].Add(s + nt);
            }
            
            grammar.P.Add(nt, new HashSet<string>());
            foreach (var s in alfa)
            {
                grammar.P[nt].Add(s);
                grammar.P[nt].Add(s + nt);
            }
        }
    }
}