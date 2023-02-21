from typing import List


class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        front, back = 0, len(nums) - 1

        while front < back:
            mid = (front + back) // 2
            if nums[mid] == nums[mid ^ 1]:
                front = mid + 1
            else:
                back = mid

        return nums[front]
