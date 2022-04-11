namespace lab5;

public static class Step5
{
    public static void ParseWord(this Grammar g)
    {
        var stack = new Stack<string>();
        stack.Push("$");
        stack.Push(g.S);
        
        var w = g.Word + "$";
        var d = "S";
        var p = 0;

        while (stack.Count > 0)
        {
            Console.WriteLine($"Word: {w}");
            Console.WriteLine($"Stack: {string.Join("", stack)}");
            Console.WriteLine($"Derivation: {d}");

            var current = w[0].ToString();
            var top = stack.Peek();

            if (current == top)
            {
                stack.Pop();
                w = w.Substring(1);
                Console.WriteLine($"\n------> Discard symbol {current}\n");
                continue;
            }
            
            var symbol = g.ParsingTable[top][current];
            
            Console.WriteLine($"\n------> Replace {top} with {symbol}\n");
            
            symbol = symbol.Replace("ε", "");
            
            d = d.Remove(d.IndexOf(top)) + symbol + d.Substring(d.IndexOf(top) + 1);
            
            stack.Pop();
            foreach (var c in symbol.Reverse())
                stack.Push(c.ToString());
        }
    }
    
}