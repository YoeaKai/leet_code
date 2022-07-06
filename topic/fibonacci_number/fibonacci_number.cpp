#include <cmath>
using namespace std;

class Solution
{
public:
	int fib(int n)
	{
		return n <= 1 ? n : fib(n - 1) + fib(n - 2);
	}

	// Refer: https://www.cuemath.com/numbers/fibonacci-sequence/
	int fib2(int n)
	{
		const double g = (1 + sqrt(5)) / 2;
		return (pow(g, n) - pow(1 - g, n)) / sqrt(5);
	}
};