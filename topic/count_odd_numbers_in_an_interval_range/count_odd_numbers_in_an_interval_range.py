class Solution:
    def countOdds(self, low: int, high: int) -> int:
        return (high - low + (1 if low % 2 == 0 else 2)) // 2
