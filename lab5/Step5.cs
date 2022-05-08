namespace lab5;

public static class Step5
{
    public static void ParseWord(this Grammar g)
    {
        var stack = new Stack<string>();
        stack.Push("$");
        stack.Push(g.S);
        
        var w = g.Word + "$";

        while (stack.Count > 0)
        {
            var current = w[0].ToString();
            var top = stack.Peek();
            var a = "";
            var str = "";
            
            if (current == top)
            {
                stack.Pop();
                w = w.Substring(1);
                a = $"-> Discard symbol {current}";
                str = $" {w} | {string.Join("", stack)} | {a}\n";
                Console.WriteLine(str);
                continue;
            }
            
            var symbol = g.ParsingTable[top][current];
            
            a = $"-> Replace {top} with {symbol}";
            
            str = $" {w} | {string.Join("", stack)} | {a}\n";
            Console.WriteLine(str);
            
            symbol = symbol.Replace("ε", "");

            stack.Pop();
            foreach (var c in symbol.Reverse())
                stack.Push(c.ToString());
        }
    }
    
}