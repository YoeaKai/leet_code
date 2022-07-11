#include <iostream>
using namespace std;

class Solution
{
public:
	bool checkOverlap(int radius, int xCenter, int yCenter, int x1, int y1, int x2, int y2)
	{
		x1 -= xCenter;
		x2 -= xCenter;
		y1 -= yCenter;
		y2 -= yCenter;
		// If x1 * x2 > 0 or y1 * y2 > 0,
		// it means the square is directly above, below, left, and right of the circle.
		int minX = x1 * x2 > 0 ? min(x1 * x1, x2 * x2) : 0, minY = y1 * y2 > 0 ? min(y1 * y1, y2 * y2) : 0;
		// The closest distance is equal to the hypotenuse of the smallest triangle.
		return minY + minX <= radius * radius;
	}
};