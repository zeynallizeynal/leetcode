class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        nums = sorted(nums)
        for i in range(len(nums)):
            print(nums)
            if nums[i] == val:
                j = i + 1
                while j < len(nums):
                    if nums[i] == nums[j]:
                        nums.pop(j)
                        j = i + 1
                    else:
                        return len(nums)

        return len(nums)


print(Solution().removeElement([0, 1, 2, 2, 3, 0, 4, 2], 2))
