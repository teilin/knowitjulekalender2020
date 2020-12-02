using System;
using System.Collections;
using System.Collections.Generic;

namespace Day02
{
    class Program
    {
        static ArrayList primes = new ArrayList();

        static void Sieve()
        {
            var n = MAX;
            var nNew = (int)Math.Sqrt(n);
            var marked = new int[n/2+500];

            for(var i = 1; i <= (nNew - 1) / 2; i++)
                for(var j = (i * (i+1)) << 1;
                                j <= n / 2; j = j + 2 * i + 1)
                                marked[j] = 1;

            primes.Add(2);

            for(var i = 1; i <= n / 2; i++)
                if(marked[i] == 0)
                    primes.Add(2*i+1);
        }
        static int BinarySearch(int left, int right, int n)
        {
            if(left <= right)
            {
                var mid = (left + right) / 2;
                if(mid == 0 || mid == primes.Count - 1)
                    return (int)primes[mid];
                
                if((int)primes[mid] == n)
                    return (int)primes[mid];

                if((int)primes[mid] < n && (int)primes[mid+1] > n)
                    return (int)primes[mid];
                if(n < (int)primes[mid])
                    return BinarySearch(left, mid-1,n);
                else
                    return BinarySearch(mid + 1, right, n);
            }
            return 0;
        }

        public static bool Check(int num)
        {
            if(num == 0) return false;
            num = Math.Abs(num);
            if(num%7 == 0) return true;
            while(num > 1)
            {
                if(num%10 == 7) return true;
                num/=10;
            }
            return false;
        }
        static bool ContainsDigit(int n, int digit)
        {
            var strN = n.ToString().ToCharArray();
            var d = digit.ToString().ToCharArray()[0];

            for(var i=0;i<strN.Length;i++) if(strN[i]==d) return true;

            return false;
        }
        static int MAX = 5433000;// 5433000

        static void Main(string[] args)
        {
            Sieve();

            var delivered = 0;
            var index = 0;

            while(index <= MAX)
            {
                var isThrownAway = ContainsDigit(index,7);

                if(!isThrownAway) 
                {
                    delivered++;
                    index++;
                }
                if(isThrownAway)
                {
                    var closestPrime = BinarySearch(0, primes.Count -1, index);
                    index += closestPrime;
                    index++;
                }
            }

            Console.WriteLine(delivered);
        }
    }
}
