class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        for i in range(0, len(nums)):
            j = i + 1
            while j < len(nums):
                if nums[i] == nums[j]:
                    nums.pop(j)
                    j = i + 1
                else:
                    break
        return nums
