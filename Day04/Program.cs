using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;

class Program
{
    static IDictionary<string, int> units = new Dictionary<string, int>
    {
        {"sukker", 2},
        {"mel", 3},
        {"melk",3},
        {"egg",1}
    };
    static IDictionary<string, int> ingredients = new Dictionary<string, int>();
    static int Calculate() 
    {
        var min = ingredients
            .Select(e => new {Ingredient = e.Key, Rel = e.Value/(double)units[e.Key]})
            .OrderBy(o => o.Rel)
            .First();
        return Convert.ToInt32(Math.Floor(min.Rel));
    }
    static void Parse(string[] delivery)
    {
        foreach(var s in delivery)
        {
            var elm = s.Split(':', StringSplitOptions.TrimEntries);
            if(ingredients.ContainsKey(elm[0])) ingredients[elm[0]] += int.Parse(elm[1].Trim());
            else ingredients.Add(elm[0], int.Parse(elm[1]));
        }
    }

    static void Main(string[] args)
    {
        var history = File
            .ReadLines(@"leveringsliste.txt")
            .Select(line => line.Split(',', StringSplitOptions.TrimEntries));

        foreach(var x in history)
            Parse(x);

        Console.WriteLine(Calculate());
    }
}
