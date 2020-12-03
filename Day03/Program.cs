using System;
using System.IO;
using System.Linq;
using System.Text;

class Program
{
    static int Rows, Columns;

    static int[] x = { -1, -1, -1, 0, 0, 1, 1, 1 };
    static int[] y = { -1, 0, 1, -1, 1, -1, 0, 1 };

    static bool Search(char[][] grid, int row, int column, string word)
    {
        if(grid[row][column]!=word[0]) return false;
        var length = word.Length;
        for(var direction=0;direction<8;direction++)
        {
            int k,rd=row+x[direction],cd=column+y[direction];
            for(k=1;k<length;k++)
            {
                if(rd>=Rows || rd<0 || cd>=Columns || cd<0) break;
                if(grid[rd][cd] != word[k]) break;
                rd += x[direction];
                cd += y[direction];
            }
            if(k==length) return true;
        }
        return false;
    }

    static bool PattarnSearch(char[][] grid, string word)
    {
        for(int row=0;row<Rows;row++)
        {
            for(var col=0;col<Columns;col++)
            {
                if(Search(grid,row,col,word)) return true;
            }
        }
        return false;
    }
    
    static void Main(string[] args)
    {
        Rows = 1000;
        Columns = 1000;

        char[][] grid = File
            .ReadLines(@"matrix.txt", Encoding.Default)
            .Select(line => line.ToCharArray())
            .ToArray();

        var wordsNotInList = File
            .ReadLines(@"wordlist.txt", Encoding.Default)
            .Where(w => !PattarnSearch(grid, w))
            .OrderBy(order => order)
            .Select(word => word)
            .ToArray();

        Console.WriteLine(string.Join(',', wordsNotInList));
    }
}