#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int numPrimeArrangements(int n) {
        vector<int> p = {2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97}; // for g++ -std=c++11
        // int const p[] = {2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97}; // for the Clang default standard: C++98
        int k=upper_bound(begin(p),p.end(),n)-begin(p);
        long long m=1e9+7,res=1;
        for(int i=2;i<=k;i++)res*=i,res%=m;
        for(int i=2;i<=n-k;i++)res*=i,res%=m;
        return res;
    }
};
