using System;
using System.Linq;
using System.IO;
using System.Text;

class Program
{
    static bool isTreeSymmetric(char[][] forrest, int startX, int startY)
    {
        var isSymmetric = true;

        for(var i=startY;i>=0;i--)
        {
            var j = 0;
            if(forrest[i][startX] == ' ' && (forrest[i][--startX]=='#' || forrest[i][++startX]=='#')) return false;
            while((-1*j)+startX >= 0 && (forrest[i][(-1*j)+startX] != ' ' || forrest[i][j+startX] != ' '))
            {
                if((-1*j)+startX >= 0 && forrest[i][(-1*j)+startX] != forrest[i][j+startX]) return false;
                j++;
            }

        }

        return isSymmetric;
    }

    static void Main(string[] args)
    {
        char[][] forrest = File
            .ReadLines(@"input.txt")
            .Select(line => line.ToCharArray())
            .ToArray();

        var forrestBottom = forrest.Length-1;
        var numSymmetricTrees = 0;
        var numTrees = 0;

        var x = 0;
        foreach(var tree in forrest[forrestBottom])
        {
            if(tree == '#')
            {
                if(isTreeSymmetric(forrest, x, forrestBottom)) numSymmetricTrees++;
                numTrees++;
            }
            x++;
        }

        Console.WriteLine(numSymmetricTrees);
        Console.WriteLine(numTrees);
    }
}
