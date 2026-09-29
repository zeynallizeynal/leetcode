class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        for i in range(0, len(nums)):
            j = i
            while j < len(nums):
                if nums[j] == val:
                    nums.pop(j)
                    j = i
                else:
                    break
        return nums


print(Solution().removeElement([3,2,2,3], 3))
